package main

type PreviewRequest struct {
	FromPoiID         int     `json:"fromPoiId"`
	AllowZeroPay      bool    `json:"allowZeroPay"`
	PoiID             int     `json:"poiId"`
	CityID            int     `json:"cityId"`
	CouponIds         []int   `json:"couponIds"`
	AddressID         int     `json:"addressId"`
	ActionSelect      int     `json:"actionSelect"`
	FromSource        int     `json:"fromSource"`
	IsUseCard         bool    `json:"isUseCard"`
	ShippingType      int     `json:"shippingType"`
	SelfLiftingMobile string  `json:"selfLiftingMobile"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
}
type ArrivalTimeWithDateRequest struct {
	AddressID    int            `json:"addressId"`
	PoiID        int            `json:"poiId"`
	StockPois    []int          `json:"stockPois"`
	ShippingType int            `json:"shippingType"`
	Longitude    float64        `json:"longitude"`
	Latitude     float64        `json:"latitude"`
	PackageTime  map[string]int `json:"packageTime"`
}

type SubmitPackageItem struct {
	EstimateTime      int64  `json:"estimateTime"`
	PackageID         int    `json:"packageId"`
	DeliverType       int    `json:"deliverType"`
	IsSpeedy          bool   `json:"isSpeedy"`
	SchemeID          int    `json:"schemeId"`
	DeliveryStartTime int64  `json:"deliveryStartTime"`
	DeliveryEndTime   int64  `json:"deliveryEndTime"`
	DateTime          int64  `json:"dateTime"`
	DeliveryLevel     int    `json:"deliveryLevel"`
	DeliveryUUID      string `json:"deliveryUuid"`
}
type SubmitRequest struct {
	AllowZeroPay       bool                `json:"allowZeroPay"`
	CityID             int                 `json:"cityId"`
	PoiID              int                 `json:"poiId"`
	PackageInfo        []SubmitPackageItem `json:"packageInfo"`
	AddressID          int                 `json:"addressId"`
	ActionSelect       int                 `json:"actionSelect"`
	CouponIds          []int64             `json:"couponIds"`
	TotalPay           float64             `json:"totalPay"`
	TimeType           int                 `json:"timeType"`
	StockPois          []int               `json:"stockPois"`
	Remark             string              `json:"remark"`
	SelfLiftingMobile  string              `json:"selfLiftingMobile"`
	ShippingType       int                 `json:"shippingType"`
	SelfLiftingAddress string              `json:"selfLiftingAddress"`
	AppID              string              `json:"appId,omitempty"`
	OpenID             string              `json:"openId,omitempty"`
}

func NewArrivalTimeWithDateRequest(conf *MeituanConfig) *ArrivalTimeWithDateRequest {
	return &ArrivalTimeWithDateRequest{
		AddressID:    conf.AddressID,
		PoiID:        conf.Poi,
		StockPois:    []int{conf.Poi},
		ShippingType: 0,
		Longitude:    conf.HomepageLng,
		Latitude:     conf.HomepageLat,
		PackageTime: map[string]int{
			"0": 0,
		},
	}
}

func NewPreviewRequest(conf *MeituanConfig) *PreviewRequest {
	return &PreviewRequest{
		FromPoiID:         0,
		AllowZeroPay:      true,
		PoiID:             conf.Poi,
		CityID:            conf.CityID,
		CouponIds:         []int{-1},
		AddressID:         conf.AddressID,
		ActionSelect:      0,
		FromSource:        0,
		IsUseCard:         false,
		ShippingType:      -1,
		SelfLiftingMobile: "",
		Longitude:         conf.HomepageLng,
		Latitude:          conf.HomepageLat,
	}
}

type PriceItem struct {
	Name              string   `json:"name"`
	OriginPrice       int      `json:"originPrice"`
	Price             int      `json:"price"`
	PriceString       string   `json:"priceString"`
	FontColor         string   `json:"fontColor"`
	FontStyle         int      `json:"fontStyle"`
	DeliveryPriceType int      `json:"deliveryPriceType"`
	MemberReduceType  int      `json:"memberReduceType"`
	CouponPriceType   int      `json:"couponPriceType"`
	IconType          int      `json:"iconType"`
	FrozenTag         int      `json:"frozenTag"`
	PlusFlag          int      `json:"plusFlag"`
	PromotionFlag     int      `json:"promotionFlag"`
	PriceRegion       int      `json:"priceRegion"`
	LabelList         []string `json:"labelList,omitempty"`
	SubTitle          string   `json:"subTitle,omitempty"`
}

type AddressInfo struct {
	AddressID      int     `json:"addressId"`
	UserID         int     `json:"userId"`
	AddressPerson  string  `json:"addressPerson"`
	Address        string  `json:"address"`
	HouseNumber    string  `json:"houseNumber"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	Mobile         string  `json:"mobile"`
	Gender         int     `json:"gender"`
	PoiID          int     `json:"poiId"`
	AddressSubType int     `json:"addressSubType"`
	CheckID        int     `json:"checkId"`
	SelectType     int     `json:"selectType"`
	AddressTipMsg  string  `json:"addressTipMsg"`
}
type Coupon struct {
	CacheID           string `json:"cacheId"`
	CanUseCouponCount int    `json:"canUseCouponCount"`
}

type PreviewPackageItem struct {
	SkuProductInfo []struct {
		PoiID               int    `json:"poiId"`
		SpuID               int    `json:"spuId"`
		SkuID               int    `json:"skuId"`
		SkuName             string `json:"skuName"`
		SubTitle            string `json:"subTitle"`
		Pic                 string `json:"pic"`
		TotalSellPrice      int    `json:"totalSellPrice"`
		TotalPromotionPrice int    `json:"totalPromotionPrice"`
		SellPrice           int    `json:"sellPrice"`
		PromotionPrice      int    `json:"promotionPrice"`
		SellUnitViewPrice   int    `json:"sellUnitViewPrice"`
		SellUnitViewName    string `json:"sellUnitViewName"`
		PromotionViewPrice  int    `json:"promotionViewPrice"`
		Unit                string `json:"unit"`
		Spec                string `json:"spec"`
		Count               int    `json:"count"`
		ViewCount           string `json:"viewCount"`
		ItemTag             int    `json:"itemTag"`
		ProcessingDetail    string `json:"processingDetail"`
		TempCount           int    `json:"tempCount"`
		FrozenTag           int    `json:"frozenTag"`
		MemberTag           int    `json:"memberTag"`
		GiftInfoTips        string `json:"giftInfoTips"`
		IsGift              bool   `json:"isGift"`
		Scatter             bool   `json:"scatter"`
	} `json:"skuProductInfo"`
	PackageID               int    `json:"packageId"`
	EstimateTimeString      string `json:"estimateTimeString"`
	DeliveryType            int    `json:"deliveryType"`
	PackageLabel            string `json:"packageLabel"`
	PackageLabelID          int    `json:"packageLabelId"`
	TotalCount              int    `json:"totalCount"`
	PackageName             string `json:"packageName"`
	EstimateTime            int    `json:"estimateTime"`
	DateTime                int    `json:"dateTime"`
	TimeDeliveryPrice       int    `json:"timeDeliveryPrice"`
	SupportSelfLifting      bool   `json:"supportSelfLifting"`
	SelfLiftingTimeString   string `json:"selfLiftingTimeString"`
	SelfLiftingTime         int    `json:"selfLiftingTime"`
	DeliveryLevel           int    `json:"deliveryLevel"`
	DeliveryStartTime       int64  `json:"deliveryStartTime"`
	DeliveryEndTime         int64  `json:"deliveryEndTime"`
	HalfDayActivityTicketID int    `json:"halfDayActivityTicketId"`
	SpeedyDelivery          struct {
		IsChooseSpeedy  bool `json:"isChooseSpeedy"`
		IsSupportSpeedy bool `json:"isSupportSpeedy"`
	} `json:"speedyDelivery"`
	SchemeID                 int    `json:"schemeId"`
	PackageWeight            string `json:"packageWeight"`
	EarliestUseTime          bool   `json:"earliestUseTime"`
	DeliveryCode             int    `json:"deliveryCode"`
	PredictArrivateTime      int64  `json:"predictArrivateTime"`
	DeliveryUUID             string `json:"deliveryUuid"`
	ShowDevliveryArrivedIcon bool   `json:"showDevliveryArrivedIcon"`
}

func NewCartRefreshRequest(PoiID int64) *CartRefreshRequest {
	return &CartRefreshRequest{
		CartOpType:   "REFRESH",
		CartOpSource: "CART",
		PoiID:        PoiID,
		ShippingType: 0,
	}
}

type CartRefreshRequest struct {
	CartOpType   string `json:"cartOpType"`
	CartOpSource string `json:"cartOpSource"`
	OpTarget     struct {
		OpTargets []interface{} `json:"opTargets"`
	} `json:"opTarget"`
	PoiID        int64 `json:"poiId"`
	ShippingType int   `json:"shippingType"`
}
