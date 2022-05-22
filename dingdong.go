package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type DingdongJob struct {
	conf    *DingDongConf
	sleepMs time.Duration
}

//接口常量信息
const (
	API_VERSION   = "9.51.0"
	APP_VERSION   = "2.86.3"
	APP_CLIENT_ID = 3
	CHANNEL       = "applet"
)

var (
	contentType            = "application/x-www-form-urlencoded"
	cartIndexUrl           = "https://maicai.api.ddxq.mobi/cart/index"
	checkOrderUrl          = "https://maicai.api.ddxq.mobi/order/checkOrder"
	addNewOrderUrl         = "https://maicai.api.ddxq.mobi/order/addNewOrder"
	getMultiReserveTimeUrl = "https://maicai.api.ddxq.mobi/order/getMultiReserveTime"
)

func NewDDJob(conf *DingDongConf, ms time.Duration) *DingdongJob {
	return &DingdongJob{conf: conf, sleepMs: ms}
}

func getCommonHeaders(conf *DingDongConf) map[string]string {
	result := make(map[string]string)
	result["ddmc-api-version"] = API_VERSION
	result["ddmc-app-client-id"] = fmt.Sprintf("%d", APP_CLIENT_ID)
	result["ddmc-city-number"] = conf.CityID
	result["ddmc-station-id"] = conf.Poi
	result["ddmc-uid"] = conf.Uid
	result["Cookie"] = fmt.Sprintf("DDXQSESSID=%s", conf.Token)
	return result
}

func getCommonParams(conf *DingDongConf) map[string]string {
	result := make(map[string]string)
	result["uid"] = conf.Uid
	result["station_id"] = conf.Poi
	result["city_number"] = conf.CityID
	result["app_version"] = APP_VERSION
	result["api_version"] = API_VERSION
	result["channel"] = CHANNEL
	result["app_client_id"] = fmt.Sprintf("%d", APP_CLIENT_ID)
	result["wx"] = "1"
	return result
}

func (d *DingdongJob) Run() {
	var err error
	//刷新购物车列表
	cartProducts, sign, err := d.getCartProducts(context.TODO())
	if err != nil {
		log.Errorf("获取购物车商品出错:%s", err.Error())
		return
	}
	log.Debugf("%+v", cartProducts)
	if len(cartProducts) == 0 {
		log.Infof("购物车为空")
		return
	}
	var orderData *OrderData
	for {
		orderData, err = d.checkOrder(&cartProducts)
		if err != nil {
			log.Errorf("checkOrder error:%s", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	//获取预约时间
	var times []ReserveTimeItem
	for {
		times, err = d.getMultiReserveTime(cartProducts)
		if err != nil {
			log.Errorf("getMultiReserveTime error:%s", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if len(times) == 0 {
		log.Info("有效预约时段为空")
		return
	}
	log.Infof("获取到%d个有效预约时间段", len(times))

	for _, reserveTime := range times {
		payment := make(map[string]interface{})
		payment["reserved_time_start"] = reserveTime.StartTimestamp
		payment["reserved_time_end"] = reserveTime.EndTimestamp
		payment["price"] = orderData.TotalMoney
		payment["freight_discount_money"] = orderData.FreightDiscountMoney
		payment["freight_money"] = orderData.FreightMoney
		payment["order_freight"] = "0.00"
		payment["parent_order_sign"] = sign
		payment["product_type"] = 1
		payment["address_id"] = d.conf.AddressID
		payment["pay_type"] = 6 //6
		//不使用VIP或折扣码
		payment["vip_money"] = ""              //
		payment["vip_buy_user_ticket_id"] = "" //
		payment["coupons_money"] = ""          //
		payment["coupons_id"] = ""             //
		code, msg, err := d.createOrder(payment, cartProducts)
		if err != nil {
			log.Errorf("createOrder error:%s", err.Error())
			return
		}
		log.Info(msg)
		//5003 商品信息有变化
		if code != 6001 {
			log.Errorf("createOrder error:%d", code)
			time.Sleep(time.Second * 10)
		}
		log.Infoln("code:%d", code)

	}

}
func (d *DingdongJob) getCartProducts(ctx context.Context) ([]NewOrderProductItem, string, error) {
	var cartData *CartData
	var err error
	//刷新购物车列表
	for {
		select {
		case <-ctx.Done():
			return nil, "", err
		default:
			cartData, err = d.cartIndex()
			if err != nil {
				log.Errorf("cartIndex:%s", err.Error())
				time.Sleep(10 * time.Second)
				continue
			}
			return cartData.NewOrderProductList, cartData.ParentOrderInfo.ParentOrderSign, nil
		}
	}
}
func (d *DingdongJob) newR() *resty.Client {
	return resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetDebug(true)
}
func (d *DingdongJob) cartIndex() (*CartData, error) {
	headers := getCommonHeaders(d.conf)
	params := getCommonParams(d.conf)
	params["is_load"] = "1"
	r := d.newR()
	response := &DingdongResponse{}
	_, err := r.R().
		SetQueryParams(params).
		SetHeaders(headers).
		SetResult(response).
		SetHeader("content-type", contentType).
		Get(cartIndexUrl)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New(response.Msg)
	}
	data := &CartData{}
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DingdongJob) checkOrder(orderProducts *[]NewOrderProductItem) (*OrderData, error) {
	headers := getCommonHeaders(d.conf)
	params := getCommonParams(d.conf)
	params["showMsg"] = "false"
	params["showData"] = "true"
	params["user_ticket_id"] = "default"
	params["freight_ticket_id"] = "default"
	params["is_use_balance"] = "0"
	params["is_buy_vip"] = "0"
	params["is_support_merge_payment"] = "1"

	for i, prod := range *orderProducts {
		prods := prod.Products
		for j, v := range prods {
			prods[j].TotalMoney = v.TotalPrice
			prods[j].TotalOriginMoney = v.TotalOriginPrice
		}
		(*orderProducts)[i].Products = prods
	}

	packages, err := json.Marshal(orderProducts)
	if err != nil {
		return nil, err
	}
	params["packages"] = string(packages)
	r := d.newR()
	response := &DingdongResponse{}
	_, err = r.R().
		SetFormData(params).
		SetHeaders(headers).
		SetResult(response).
		SetHeader("content-type", contentType).
		Post(checkOrderUrl)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New(response.Msg)
	}
	data := &checkOrderData{}
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		return nil, err
	}
	return &data.Order, nil
}

func (d *DingdongJob) getMultiReserveTime(list []NewOrderProductItem) ([]ReserveTimeItem, error) {
	headers := getCommonHeaders(d.conf)
	params := getCommonParams(d.conf)
	params["products"] = "[[{}]]"
	response := &DingdongResponse{}
	r := d.newR()
	_, err := r.R().
		SetFormData(params).
		SetHeaders(headers).
		SetResult(response).
		SetHeader("content-type", contentType).
		Post(getMultiReserveTimeUrl)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New(response.Msg)
	}
	data := []getMultiReserveTimeItem{}
	err = json.Unmarshal(response.Data, &data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("获取预约出错")
	}
	if len(data[0].Time) == 0 {
		return nil, errors.New("获取预约出错:time为空")

	}
	times := data[0].Time[0].Times
	var result []ReserveTimeItem
	for _, v := range times {
		if v.FullFlag {
			continue
		}
		result = append(result, v)
	}
	return result, nil
}

func (d *DingdongJob) createOrder(payment map[string]interface{}, packages []NewOrderProductItem) (int, string, error) {
	headers := getCommonHeaders(d.conf)
	params := getCommonParams(d.conf)
	params["showMsg"] = "false"
	params["showData"] = "true"
	params["ab_config"] = `{"key_onion":"C"}`
	params["address_id"] = d.conf.AddressID
	for i := range packages {
		packages[i].ReservedTimeStart = payment["reserved_time_start"].(int64)
		packages[i].ReservedTimeEnd = payment["reserved_time_end"].(int64)
		packages[i].SoonArrival = ""
		packages[i].FirstSelectedBigTime = 0
	}
	packageOrder := make(map[string]interface{})
	packageOrder["payment_order"] = payment
	packageOrder["packages"] = packages
	temp, err := json.Marshal(packageOrder)
	if err != nil {
		return 0, "", err
	}
	params["package_order"] = string(temp)
	r := resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetDebug(true)
	response := &DingdongResponse{}
	_, err = r.R().
		SetFormData(params).
		SetHeaders(headers).
		SetResult(response).
		SetHeader("content-type", contentType).
		Post(addNewOrderUrl)
	if err != nil {
		return 0, "", err
	}
	return response.Code, response.Msg, nil
}
