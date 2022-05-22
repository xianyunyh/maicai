package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	wechatUA       = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"
	previewUrl     = "https://mall.meituan.com/api/c/mallorder/preview"
	arrivalTimeTpl = "https://mall.meituan.com/api/c/mallorder/%d/arrivalTimeWithDate"
	submitUrl      = "https://mall.meituan.com/api/c/mallorder/submit"
	addListUrl     = "https://mall.meituan.com/api/c/malluser/address/list"
	refreshCartUrl = "https://mall.meituan.com/api/c/malluser/cart/v2/items"
	timeFormat     = "2006-01-02 15:04:05"
)

var (
	uniqId = uuid.NewString()
)

func commanGetParams(conf *MeituanConfig) url.Values {
	values := url.Values{}
	values.Set("utm_medium", "wxapp")
	values.Set("platform", "android")
	values.Set("brand", "xiaoxiangmaicai")
	values.Set("tenantId", "1")
	values.Set("utm_term", conf.UtmTerm)
	values.Set("address_id", strconv.Itoa(conf.AddressID))
	values.Set("poi", strconv.Itoa(conf.Poi))
	values.Set("stockPois", strconv.Itoa(conf.Poi))
	values.Set("homepageLng", fmt.Sprintf("%f", conf.HomepageLng))
	values.Set("homepageLat", fmt.Sprintf("%f", conf.HomepageLat))
	if conf.Uid > 0 {
		values.Set("userid", fmt.Sprintf("%d", conf.Uid))
	}
	return values
}

func AddressList(uiqId string, conf *MeituanConfig) ([]AddressItem, error) {
	client := &http.Client{}
	params := url.Values{}
	params.Set("utm_medium", "wxapp")
	params.Set("platform", "android")
	params.Set("brand", "xiaoxiangmaicai")
	params.Set("tenantId", "1")
	params.Set("uuid", uiqId)
	reqUrl := addListUrl + "?" + params.Encode()
	req, _ := http.NewRequest("GET", reqUrl, nil)
	head := http.Header{}
	head.Set("t", conf.Token)
	head.Set("User-Agent", wechatUA)
	head.Set("Content-Type", "application/json")
	req.Header = head
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	prevResp := &Response{}
	err = json.Unmarshal(data, prevResp)
	if err != nil {
		return nil, err
	}
	if prevResp.Code != 0 {
		return nil, errors.New(prevResp.Error.Msg)
	}
	respData := struct {
		AddressList []AddressItem `json:"addressList"`
	}{}
	err = json.Unmarshal(prevResp.Data, &respData)
	if err != nil {
		return nil, err
	}
	return respData.AddressList, nil
}

type optitons struct {
	conf  *MeituanConfig
	uiqId string
}
type Option = func(o *optitons)

func WithConf(c *MeituanConfig) Option {
	return func(o *optitons) {
		o.conf = c
	}
}

func WithUiqId(id string) Option {
	return func(o *optitons) {
		o.uiqId = id
	}
}

func MeituanReq(uri string, method string, body interface{}, opts ...Option) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var (
		reqBody []byte
		err     error
	)
	if body != nil {
		reqBody, err = json.Marshal(body)
	}
	if err != nil {
		return nil, err
	}
	args := &optitons{
		conf: &MeituanConfig{},
	}
	for _, fn := range opts {
		fn(args)
	}
	start := time.Now()
	params := commanGetParams(args.conf)
	params.Set("uuid", args.uiqId)
	params.Set("xuuid", args.uiqId)
	reqUrl := uri + "?" + params.Encode()
	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	head := http.Header{}
	head.Set("t", args.conf.Token)
	head.Set("User-Agent", wechatUA)
	head.Set("Content-Type", "application/json")
	req.Header = head
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http_code=%d", resp.StatusCode)
	}
	end := time.Now()
	log.Debugf("api=[%s] 耗时:%d ms", reqUrl, end.Sub(start).Milliseconds())
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	temp := &Response{}
	log.WithField("req_url", reqUrl).
		WithField("request", string(reqBody)).
		Debugf("response=%s", data)
	err = json.Unmarshal(data, temp)
	if err != nil {
		return nil, err
	}
	if temp.Code != 0 {
		log.Errorf("error_code=[%d],msg=[%s]", temp.Code, temp.Error.Msg)
		return nil, ReponseError{code: temp.Code, message: temp.Error.Msg}
	}
	return temp.Data, nil
}

//预览购物车
func MTPreview(uiqId string, request *PreviewRequest, conf *MeituanConfig) (*PreviewData, error) {

	data, err := MeituanReq(previewUrl, "POST", request, WithConf(conf), WithUiqId(uiqId))
	if err != nil {
		return nil, err
	}
	respData := &PreviewData{}
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

//刷新购物车

func MTRefreshCart(uiqId string, request *CartRefreshRequest, conf *MeituanConfig) (*CartRefreshData, error) {
	data, err := MeituanReq(refreshCartUrl, "POST", request, WithConf(conf), WithUiqId(uiqId))
	if err != nil {
		return nil, err
	}
	respData := &CartRefreshData{}
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

//获取配送
func MTArrivalTimeWithDate(uiqId string, conf *MeituanConfig) (*ArrivalTimeWithDateData, error) {
	request := NewArrivalTimeWithDateRequest(conf)
	reqUrl := fmt.Sprintf(arrivalTimeTpl, conf.Poi)
	data, err := MeituanReq(reqUrl, "POST", request, WithConf(conf), WithUiqId(uniqId))
	if err != nil {
		return nil, err
	}
	respData := &ArrivalTimeWithDateData{}
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

//提交订单
func MTSubmit(uiqId string, conf *MeituanConfig, request *SubmitRequest) (*SubmitResponse, error) {
	data, err := MeituanReq(submitUrl, "POST", request, WithUiqId(uiqId), WithConf(conf))
	if err != nil {
		return nil, err
	}
	respData := &SubmitResponse{}
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, err
	}
	return respData, nil

}

type MeiTuanJob struct {
	conf    *MeituanConfig
	notify  Notifyer
	SleepMs time.Duration
}

func (m *MeiTuanJob) GetPreviewOrder(ctx context.Context, res chan<- *PreviewData) {
	for {
		select {
		case <-ctx.Done():
			log.Errorf("超过脚本最长运行时间")
			res <- nil
			return
		default:
			prevReq := NewPreviewRequest(m.conf)
			temp, e := MTPreview(uniqId, prevReq, m.conf)
			if e == nil {
				log.Infof("生成预付订单成功")
				res <- temp
				return
			}
			log.Errorf("预付订单生成失败:%s, 正在重试", e.Error())
		}

		sleepMs := time.Duration(m.SleepMs)
		log.Infof("停顿%dms", sleepMs)
		time.Sleep(sleepMs * time.Millisecond)
		m.refreshCart()
	}
}

func (m *MeiTuanJob) GetArrivalTimeData(ctx context.Context, res chan<- *ArrivalTimeWithDateData) {
	for {
		select {
		case <-ctx.Done():
			log.Errorf("超过最长时间:[%v]", ctx.Err())
			res <- nil
			return
		default:
			temp, e := MTArrivalTimeWithDate(uniqId, m.conf)
			if e == nil {
				log.Info("成功获取有效配送时间")
				res <- temp
				return
			}
			log.Errorf("获取配送时间遇到错误：%s", e.Error())
		}
		sleepMs := time.Duration(m.SleepMs)
		log.Infof("停顿%dms", sleepMs)
		time.Sleep(sleepMs * time.Millisecond)
	}
}

func (m *MeiTuanJob) createOrder(timeItems ArrivalTimePackageItem, total float64) error {
	var err error
	var foundIdx int
	for _, aTitem := range timeItems.ArrivalTimeList {
		if aTitem.Disable {
			log.Infof("%s已约满", aTitem.TimeIntervals)
			continue
		}
		foundIdx = foundIdx + 1
		// 跳过第一个预约时间段
		if foundIdx == 1 {
			continue
		}
		packageItem := SubmitPackageItem{
			EstimateTime:      (aTitem.DeliveryStartTime + aTitem.DeliveryEndTime) / 2,
			PackageID:         0,
			DeliverType:       1,
			IsSpeedy:          false,
			SchemeID:          -1,
			DeliveryStartTime: aTitem.DeliveryStartTime,
			DeliveryEndTime:   aTitem.DeliveryEndTime,
			DateTime:          aTitem.DateTime,
			DeliveryLevel:     aTitem.DeliveryLevel,
			DeliveryUUID:      timeItems.DeliveryUUID,
		}
		subReq := &SubmitRequest{
			AllowZeroPay: true,
			CityID:       m.conf.CityID,
			PoiID:        m.conf.Poi,
			PackageInfo:  []SubmitPackageItem{packageItem},
			AddressID:    m.conf.AddressID,
			ActionSelect: 0,
			TotalPay:     total,
			StockPois: []int{
				m.conf.Poi,
			},
			ShippingType: 0,
		}
		log.Info("正在提交订单......")
		_, err = MTSubmit(uniqId, m.conf, subReq)

		if err != nil {
			if strings.Contains(err.Error(), "订单未支付成功") {
				log.Infof("配送时间:%s", aTitem.TimeIntervals)
				log.Info("订单生成成功，请前往订单页支付")
				return nil
			}
			log.Errorf("创建订单遇到错误:%s", err.Error())
		}
	}
	return err
}

func (m *MeiTuanJob) refreshCart() error {
	req := NewCartRefreshRequest(int64(m.conf.Poi))
	resp, err := MTRefreshCart(uniqId, req, m.conf)
	if err != nil {
		log.Errorf("刷新购物遇到错误:%s", err.Error())
		return err
	}
	log.Infof("刷新购物车成功，购物车总计商品:%d个：总价:%d", resp.TotalItemCounts, resp.TotalAmount)
	return nil
}
func (m *MeiTuanJob) Run() {
	log.Infof("开始运行：%s", time.Now().Format(timeFormat))
	var err error
	uniqId = uuid.NewString()
	previewResult := make(chan *PreviewData, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	m.GetPreviewOrder(ctx, previewResult)
	prevResp := <-previewResult
	if prevResp == nil {
		log.Error("生成订单出错")
		return
	}
	timeResult := make(chan *ArrivalTimeWithDateData, 1)
	m.GetArrivalTimeData(ctx, timeResult)
	timeR := <-timeResult
	if timeR == nil {
		log.Error("获取配送时间失败:")
		return
	}
	if len(timeR.PackageInfo) == 0 {
		log.Errorf("未找到有效配送时间段:[%v]", timeR)
		return
	}
	timeItems := timeR.PackageInfo[0]

	err = m.createOrder(timeItems, prevResp.TotalPay)

	if err != nil {
		log.Errorf("创建订单结束:%s", err.Error())
	}
	log.Infof("运行结束：%s", time.Now().Format(timeFormat))
	m.notify.Send()
}
