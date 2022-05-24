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
	conf    *DingDongConfig
	debug   bool
	sleepMs time.Duration
	finish  chan bool
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

func NewDDJob(conf *DingDongConfig, ms time.Duration, debug bool) *DingdongJob {
	if ms <= 0 {
		ms = 500
	}
	return &DingdongJob{conf: conf, sleepMs: ms, debug: debug}
}

func getCommonHeaders(conf *DingDongConfig) map[string]string {
	result := make(map[string]string)
	result["ddmc-api-version"] = API_VERSION
	result["ddmc-app-client-id"] = fmt.Sprintf("%d", APP_CLIENT_ID)
	result["ddmc-city-number"] = conf.CityID
	result["ddmc-station-id"] = conf.Poi
	result["ddmc-uid"] = conf.Uid
	result["Cookie"] = fmt.Sprintf("DDXQSESSID=%s", conf.Token)
	return result
}

func getCommonParams(conf *DingDongConfig) map[string]string {
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

func overTime(start time.Time, max time.Duration) bool {
	now := time.Now()
	if now.Sub(start) > max {
		log.Info("已经超过最长时间:%d", max/time.Minute)
		return true
	}
	return false
}

func (d *DingdongJob) Run() {
	startTime := time.Now()
	maxRunTime := time.Minute * 15
	log.Infof("【叮咚】脚本运行开始时间:%s", startTime.Format("2006-01-02 15:04:05"))
	var err error
	//刷新购物车列表
	cartProducts, sign, err := d.getCartProducts(context.TODO())
	if err != nil {
		log.Errorf("【叮咚】获取购物车商品出错:%s", err.Error())
		return
	}

	if len(cartProducts) == 0 {
		log.Infof("【叮咚】购物车为空")
		return
	}
	//校验订单
	var orderData *OrderData
	for {
		if overTime(startTime, maxRunTime) {
			return
		}
		orderData, err = d.checkOrder(&cartProducts)
		if err != nil {
			log.Errorf("【叮咚】checkOrder error:%s", err.Error())
			time.Sleep(time.Millisecond * d.sleepMs)
			continue
		}
		break
	}
	//获取预约时间
	var times []ReserveTimeItem
	for {
		if overTime(startTime, maxRunTime) {
			return
		}
		times, err = d.getMultiReserveTime(cartProducts)
		if err != nil {
			log.Errorf("【叮咚】getMultiReserveTime error:%s", err.Error())
			time.Sleep(time.Millisecond * d.sleepMs)
			continue
		}
		break
	}
	if len(times) == 0 {
		log.Info("【叮咚】有效预约时段为空")
		return
	}
	log.Infof("【叮咚】获取到%d个有效预约时间段", len(times))
loop:
	if overTime(startTime, maxRunTime) {
		log.Info("【叮咚】脚本已经超过最长时间")
		return
	}
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
		payment["pay_type"] = 6 //小程序支付
		//不使用VIP或折扣码
		payment["vip_money"] = ""              //
		payment["vip_buy_user_ticket_id"] = "" //
		payment["coupons_money"] = ""          //
		payment["coupons_id"] = ""             //
		log.Infof("【叮咚】尝试时段:%s", reserveTime.SelectMsg)
		code, msg, err := d.createOrder(payment, cartProducts)
		//错误直接返回
		if err != nil {
			log.Errorf("createOrder error:%s", err.Error())
			return
		}
		//6001 支付参数错误
		if code == 6001 {
			log.Info(msg)
			log.Infof("【叮咚】下单成功✔️:%s😃😃😃😃😃😃😃", reserveTime.SelectMsg)
			return
		}
		log.Errorf("【叮咚】创建订单 error:%d msg:%s", code, msg)
		time.Sleep(2 * time.Second)
	}
	log.Infof("【叮咚】开始重试门店预约时间")

	goto loop
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
				log.Errorf("【叮咚】cartIndex:%s", err.Error())
				time.Sleep(time.Millisecond * d.sleepMs)
				continue
			}
			return cartData.NewOrderProductList, cartData.ParentOrderInfo.ParentOrderSign, nil
		}
	}
}
func (d *DingdongJob) newR() *resty.Request {
	return resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetDebug(d.debug).R()
}

func (d *DingdongJob) cartIndex() (*CartData, error) {
	headers := getCommonHeaders(d.conf)
	params := getCommonParams(d.conf)
	params["is_load"] = "1"
	r := d.newR()
	response := &DingdongResponse{}
	_, err := r.
		SetQueryParams(params).
		SetHeaders(headers).
		SetResult(response).
		SetHeader("content-type", contentType).
		SetHeader("User-agent", wechatUA).
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
	_, err = r.
		SetFormData(params).
		SetHeaders(headers).
		SetResult(response).
		SetHeader("content-type", contentType).
		// SetHeader("User-agent", wechatUA).
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
	_, err := r.
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
		return nil, errors.New("【叮咚】获取预约出错")
	}
	if len(data[0].Time) == 0 {
		return nil, errors.New("获取预约出错:time为空")

	}
	times := data[0].Time[0].Times
	var result []ReserveTimeItem
	for _, v := range times {
		//跳过约满
		if v.FullFlag {
			continue
		}
		//跳过type !=1
		if v.Type != 1 {
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
	r := d.newR()
	response := &DingdongResponse{}
	_, err = r.
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
