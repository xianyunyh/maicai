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
	log.Debugf("????????????api=[%s] ??????:%d ms", reqUrl, end.Sub(start).Milliseconds())
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
		log.Errorf("????????????error_code=[%d],msg=[%s]", temp.Code, temp.Error.Msg)
		return nil, ReponseError{code: temp.Code, message: temp.Error.Msg}
	}
	return temp.Data, nil
}

//???????????????
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

//???????????????

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

//????????????
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

//????????????
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

type Job interface {
	Run()
}
type MeiTuanJob struct {
	conf    *MeituanConfig
	notify  Notifyer
	SleepMs time.Duration
	finish  chan bool
}

func (m *MeiTuanJob) GetPreviewOrder(ctx context.Context, res chan<- *PreviewData) {
	for {
		select {
		case <-ctx.Done():
			log.Errorf("??????????????????????????????????????????")
			res <- nil
			return
		default:
			prevReq := NewPreviewRequest(m.conf)
			temp, e := MTPreview(uniqId, prevReq, m.conf)
			if e == nil {
				log.Infof("????????????????????????????????????")
				res <- temp
				return
			}
			log.Errorf("????????????????????????????????????:%s, ????????????", e.Error())
		}

		sleepMs := time.Duration(m.SleepMs)
		log.Infof("??????????????????????????????%dms", sleepMs)
		time.Sleep(sleepMs * time.Millisecond)
		count, _, err := m.refreshCart()
		if err != nil {
			log.Errorf("???????????????????????????????????????:%s", err.Error())
			continue
		}
		if count == 0 {
			log.Infof("???????????????????????????")
			res <- nil
			return
		}
	}
}

func (m *MeiTuanJob) GetArrivalTimeData(ctx context.Context, res chan<- *ArrivalTimeWithDateData) {
	for {
		select {
		case <-ctx.Done():
			log.Errorf("??????????????????????????????:[%v]", ctx.Err())
			res <- nil
			return
		default:
			temp, e := MTArrivalTimeWithDate(uniqId, m.conf)
			if e == nil {
				log.Info("??????????????????????????????????????????")
				res <- temp
				return
			}
			log.Errorf("?????????????????????????????????????????????%s", e.Error())
		}
		sleepMs := time.Duration(m.SleepMs)
		log.Infof("??????????????????%dms", sleepMs)
		time.Sleep(sleepMs * time.Millisecond)
	}
}

func (m *MeiTuanJob) createOrder(timeItems ArrivalTimePackageItem, total float64) error {
	var err error
	var foundIdx int
	for _, aTitem := range timeItems.ArrivalTimeList {
		if aTitem.Disable {
			log.Infof("????????????%s?????????", aTitem.TimeIntervals)
			continue
		}
		foundIdx = foundIdx + 1
		// ??????????????????????????????
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
		log.Info("??????????????????????????????......")
		_, err = MTSubmit(uniqId, m.conf, subReq)

		if err != nil {
			if strings.Contains(err.Error(), "?????????????????????") {
				log.Infof("????????????????????????:%s", aTitem.TimeIntervals)
				log.Info("???????????????????????????????????????????????????????????????????????????????????????")
				return nil
			}
			log.Errorf("????????????????????????:%s", err.Error())
		}
	}
	return err
}

func (m *MeiTuanJob) refreshCart() (int, int, error) {
	req := NewCartRefreshRequest(int64(m.conf.Poi))
	resp, err := MTRefreshCart(uniqId, req, m.conf)
	if err != nil {
		log.Errorf("????????????????????????:%s", err.Error())
		return 0, 0, err
	}
	log.Infof("?????????????????????????????????????????????:%d????????????:%d", resp.TotalItemCounts, resp.TotalAmount)
	return resp.TotalItemCounts, resp.TotalAmount, nil
}
func (m *MeiTuanJob) Run() {
	log.Infof("?????????????????????????????????%s", time.Now().Format(timeFormat))
	var err error
	uniqId = uuid.NewString()
	previewResult := make(chan *PreviewData, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	m.GetPreviewOrder(ctx, previewResult)
	prevResp := <-previewResult
	if prevResp == nil {
		log.Error("??????????????????????????????")
		return
	}
	timeResult := make(chan *ArrivalTimeWithDateData, 1)
	m.GetArrivalTimeData(ctx, timeResult)
	timeR := <-timeResult
	if timeR == nil {
		log.Error("????????????????????????????????????:")
		return
	}
	if len(timeR.PackageInfo) == 0 {
		log.Errorf("??????????????????????????????????????????:[%v]", timeR)
		return
	}
	timeItems := timeR.PackageInfo[0]

	err = m.createOrder(timeItems, prevResp.TotalPay)

	if err != nil {
		log.Errorf("??????????????????????????????:%s", err.Error())
	}
	log.Infof("???????????????????????????%s", time.Now().Format(timeFormat))
}
