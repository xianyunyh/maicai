package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code  int             `json:"code"`
	Data  json.RawMessage `json:"data"`
	Error *ResponseError  `json:"error"`
}

type ResponseError struct {
	Msg string `json:"msg"`
}

type SKUItem struct {
	StyleType      int `json:"styleType"`
	PromotionLabel struct {
		Text string `json:"text"`
	} `json:"promotionLabel"`
	PriceMarkLabel struct {
		FinalPriceColor string `json:"finalPriceColor"`
		Display         bool   `json:"display"`
	} `json:"priceMarkLabel"`
}

type PreviewData struct {
	PriceInfo            []PriceItem        `json:"priceInfo"`
	CouponReducePrice    int                `json:"couponReducePrice"`
	AddressInfo          AddressInfo        `json:"addressInfo"`
	TotalPay             float64            `json:"totalPay"`
	PriceThresholdStatus int                `json:"priceThresholdStatus"`
	DiffPriceThreshold   int                `json:"diffPriceThreshold"`
	ItemStyleData        map[string]SKUItem `json:"itemStyleData"`
	DeliveryType         int                `json:"deliveryType"`
	Coupons              struct {
		DefaultChoiceCouponIDSet []int64 `json:"defaultChoiceCouponIdSet"`
		BestCouponLabel          string  `json:"bestCouponLabel"`
		ChoiceCouponIDSet        []int64 `json:"choiceCouponIdSet"`
		CacheID                  string  `json:"cacheId"`
		CanUseCouponCount        int     `json:"canUseCouponCount"`
	} `json:"coupons"`
	PackageInfo []struct {
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
			ItemStyleKey        string `json:"itemStyleKey,omitempty"`
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
		EstimateTime            int64  `json:"estimateTime"`
		DateTime                int    `json:"dateTime"`
		TimeDeliveryPrice       int    `json:"timeDeliveryPrice"`
		SupportSelfLifting      bool   `json:"supportSelfLifting"`
		SelfLiftingTimeString   string `json:"selfLiftingTimeString"`
		SelfLiftingTime         int64  `json:"selfLiftingTime"`
		DeliveryLevel           int    `json:"deliveryLevel"`
		DeliveryStartTime       int64  `json:"deliveryStartTime"`
		DeliveryEndTime         int64  `json:"deliveryEndTime"`
		HalfDayActivityTicketID int    `json:"halfDayActivityTicketId"`
		SpeedyDelivery          struct {
			IsSupportSpeedy bool `json:"isSupportSpeedy"`
			IsChooseSpeedy  bool `json:"isChooseSpeedy"`
		} `json:"speedyDelivery"`
		SchemeID                 int    `json:"schemeId"`
		PackageWeight            string `json:"packageWeight"`
		EarliestUseTime          bool   `json:"earliestUseTime"`
		DeliveryCode             int    `json:"deliveryCode"`
		PredictArrivateTime      int64  `json:"predictArrivateTime"`
		DeliveryUUID             string `json:"deliveryUuid"`
		ShowDevliveryArrivedIcon bool   `json:"showDevliveryArrivedIcon"`
	} `json:"packageInfo"`
	DateTime          int64 `json:"dateTime"`
	DeliveryLabelNote struct {
		PopUpTitle        string        `json:"popUpTitle"`
		PopUpContent      string        `json:"popUpContent"`
		DeliveryTimeTypes []interface{} `json:"deliveryTimeTypes"`
	} `json:"deliveryLabelNote"`
	PromotionPrice         int `json:"promotionPrice"`
	AddressSelect          int `json:"addressSelect"`
	LocationDeliveryMethod int `json:"locationDeliveryMethod"`
	CouponAssign           struct {
		ConfigIds []interface{} `json:"configIds"`
		ID        int           `json:"id"`
		IDStr     string        `json:"idStr"`
	} `json:"couponAssign"`
	DiscountInfo struct {
		TotalDiscount     int `json:"totalDiscount"`
		PromotionDiscount int `json:"promotionDiscount"`
		DiscountDetails   []struct {
			Name          string `json:"name"`
			PriceString   string `json:"priceString"`
			PromotionFlag int    `json:"promotionFlag"`
		} `json:"discountDetails"`
	} `json:"discountInfo"`
	SelfLifting struct {
		SupportSelfLifting      bool    `json:"supportSelfLifting"`
		SelfLiftingAddress      string  `json:"selfLiftingAddress"`
		SelfLongitude           float64 `json:"selfLongitude"`
		SelfLatitude            float64 `json:"selfLatitude"`
		SelfLiftingMobile       string  `json:"selfLiftingMobile"`
		SelfLiftingDeliveryType int     `json:"selfLiftingDeliveryType"`
		SelfLiftingAddressID    int     `json:"selfLiftingAddressId"`
		SelfLiftingAddressName  string  `json:"selfLiftingAddressName"`
		SelfDistance            string  `json:"selfDistance"`
	} `json:"selfLifting"`
	DeliveryPrice int `json:"deliveryPrice"`
	UseSwitch     struct {
		Member struct {
			Title      string `json:"title"`
			Subtitle   string `json:"subtitle"`
			Note       string `json:"note"`
			Open       bool   `json:"open"`
			RubbishMap struct {
				Xid                string `json:"xid"`
				SalePrice          string `json:"salePrice"`
				MemberStatusPicURL string `json:"memberStatusPicUrl"`
				ReducePrice        string `json:"reducePrice"`
				Origin             string `json:"origin"`
				DetailURL          string `json:"detailUrl"`
			} `json:"rubbishMap"`
			CornerTitle string `json:"cornerTitle"`
		} `json:"member"`
	} `json:"useSwitch"`
	RuleInfo             []interface{} `json:"ruleInfo"`
	SelectDeliveryTab    int           `json:"selectDeliveryTab"`
	SupportDeliveryLevel bool          `json:"supportDeliveryLevel"`
	SpecialTimeTipMsg    string        `json:"specialTimeTipMsg"`
	PreviewPopup         struct {
		Type       int    `json:"type"`
		Title      string `json:"title"`
		ButtonList []struct {
			Desc string `json:"desc"`
			Code int    `json:"code"`
		} `json:"buttonList"`
		AutoSelected    int    `json:"autoSelected"`
		AbStrategy      string `json:"abStrategy"`
		PopupItemVOList []struct {
			SpuID             int    `json:"spuId"`
			SkuID             int    `json:"skuId"`
			SkuName           string `json:"skuName"`
			SubTitle          string `json:"subTitle"`
			Pic               string `json:"pic"`
			Spec              string `json:"spec"`
			ViewCount         string `json:"viewCount"`
			SellPrice         int    `json:"sellPrice"`
			PromotionPrice    int    `json:"promotionPrice"`
			Unit              string `json:"unit"`
			ViewUnit          string `json:"viewUnit"`
			SellUnitViewName  string `json:"sellUnitViewName"`
			Scatter           bool   `json:"scatter"`
			PromotionID       int    `json:"promotionId"`
			PromotionType     int    `json:"promotionType"`
			PromotionRuleDesc string `json:"promotionRuleDesc"`
		} `json:"popupItemVOList"`
		CancelCount    int `json:"cancelCount"`
		Interval       int `json:"interval"`
		PromotionLimit struct {
			Num5186426 int `json:"5186426"`
			Num5189631 int `json:"5189631"`
		} `json:"promotionLimit"`
	} `json:"previewPopup"`
	DeliveryRuleMsg string `json:"deliveryRuleMsg"`
	DeliveryRuleURL string `json:"deliveryRuleUrl"`
	PayInfo         struct {
		OneClickPayInfo struct {
			OpenStatus          bool   `json:"openStatus"`
			UsableStatus        bool   `json:"usableStatus"`
			DefaultPayType      int    `json:"defaultPayType"`
			CommonPaySubTitle   string `json:"commonPaySubTitle"`
			GuideMark           bool   `json:"guideMark"`
			DefaultSwitchStatus bool   `json:"defaultSwitchStatus"`
			GuideMessage        string `json:"guideMessage"`
			InstructionURL      string `json:"instructionUrl"`
			SubGuideURL         string `json:"subGuideUrl"`
			SubGuideMessage     string `json:"subGuideMessage"`
			SerialCode          string `json:"serialCode"`
		} `json:"oneClickPayInfo"`
		PayTypeButtonText struct {
			Num1 string `json:"1"`
			Num2 string `json:"2"`
		} `json:"payTypeButtonText"`
	} `json:"payInfo"`
	AbTest struct {
		Num2 int `json:"2"`
	} `json:"abTest"`
	IsShowTimeModule bool `json:"isShowTimeModule"`
}

type ArrivalTimeItem struct {
	TimeIntervals       string `json:"timeIntervals"`
	CenterTime          int64  `json:"centerTime"`
	Presell             bool   `json:"presell"`
	TimeType            int    `json:"timeType"`
	Disable             bool   `json:"disable"`
	DeliveryLevel       int    `json:"deliveryLevel"`
	DeliveryStartTime   int64  `json:"deliveryStartTime"`
	DeliveryEndTime     int64  `json:"deliveryEndTime"`
	SchemeID            int    `json:"schemeId"`
	DateTime            int64  `json:"dateTime"`
	EarliestUseTime     bool   `json:"earliestUseTime"`
	PredictArrivateTime int    `json:"predictArrivateTime"`
	IsSpeedy            bool   `json:"isSpeedy"`
}

type ArrivalTimePackageItem struct {
	PackageID       int               `json:"packageId"`
	ArrivalTimeList []ArrivalTimeItem `json:"arrivalTimeList"`
	ArrivalDateList []struct {
		Date     string `json:"date"`
		DateTime int64  `json:"dateTime"`
	} `json:"arrivalDateList"`
	SelfLiftingArrivalTimeList []interface{} `json:"selfLiftingArrivalTimeList"`
	SelfLiftingArrivalDateList []interface{} `json:"selfLiftingArrivalDateList"`
	DeliveryUUID               string        `json:"deliveryUuid"`
}
type ArrivalTimeWithDateData struct {
	PackageInfo        []ArrivalTimePackageItem `json:"packageInfo"`
	ArrivalTimeShowMsg string                   `json:"arrivalTimeShowMsg"`
	SpecialTimeTipMsg  string                   `json:"specialTimeTipMsg"`
	NotRefreshPreview  bool                     `json:"notRefreshPreview"`
}
type SubmitResponse struct {
	OrderId string `json:"orderId"`
}

type AddressItem struct {
	AddressID              int64   `json:"addressId"`
	UserID                 int64   `json:"userId"`
	AddressPerson          string  `json:"addressPerson"`
	Address                string  `json:"address"`
	HouseNumber            string  `json:"houseNumber"`
	Longitude              float64 `json:"longitude"`
	Latitude               float64 `json:"latitude"`
	Mobile                 string  `json:"mobile"`
	Gender                 int     `json:"gender"`
	AddressSubType         int     `json:"addressSubType"`
	CheckID                int     `json:"checkId"`
	SelectType             int     `json:"selectType"`
	Mid                    string  `json:"mid"`
	CheckOutAreaResultText string  `json:"checkOutAreaResultText"`
	IsDispatch             int     `json:"isDispatch"`
}

type CartItem struct {
	SortID     int64 `json:"sortId"`
	EntryType  int   `json:"entryType"`
	TotalPrice int   `json:"totalPrice"`
	TotalCount int   `json:"totalCount"`
	BaseItems  []struct {
		SortID        int64  `json:"sortId"`
		EntryType     int    `json:"entryType"`
		SkuID         int    `json:"skuId"`
		SpuID         int    `json:"spuId"`
		Title         string `json:"title"`
		SubTitle      string `json:"subTitle"`
		Pic           string `json:"pic"`
		DetailPagePic string `json:"detailPagePic"`
		Spec          string `json:"spec"`
		Count         int    `json:"count"`
		Selected      bool   `json:"selected"`
		SellerPrice   int    `json:"sellerPrice"`
		Unit          string `json:"unit"`
		PoiID         int    `json:"poiId"`
		StockPois     []int  `json:"stockPois"`
		CartItemStyle struct {
			StyleType       int `json:"styleType"`
			SimilarCheckURL struct {
				Text string `json:"text"`
				URL  string `json:"url"`
			} `json:"similarCheckUrl"`
		} `json:"cartItemStyle"`
		PackageType         int `json:"packageType"`
		SellPriceUnitInfoV2 struct {
			SellUnitViewPrice int    `json:"sellUnitViewPrice"`
			SellUnitViewName  string `json:"sellUnitViewName"`
			ViewQuantity      string `json:"viewQuantity"`
			ViewUnitName      string `json:"viewUnitName"`
			ShowMemberPrice   bool   `json:"showMemberPrice"`
		} `json:"sellPriceUnitInfoV2"`
		CanReduce              bool   `json:"canReduce"`
		StepLength             string `json:"stepLength"`
		ShowSpecPriceExplainIs bool   `json:"showSpecPriceExplainIs"`
		CanSelected            bool   `json:"canSelected"`
		CanEdit                bool   `json:"canEdit"`
		SubType                int    `json:"subType"`
		EstimatedReceiptPrice  string `json:"estimatedReceiptPrice"`
	} `json:"baseItems"`
	AllSelected   bool   `json:"allSelected"`
	GiftShowTop   bool   `json:"giftShowTop"`
	GiftCanDelete bool   `json:"giftCanDelete"`
	HeapType      string `json:"heapType"`
}

type ActivityItem struct {
	EntryType   int           `json:"entryType"`
	TotalPrice  int           `json:"totalPrice"`
	TotalCount  int           `json:"totalCount"`
	BaseItems   []interface{} `json:"baseItems"`
	AllSelected bool          `json:"allSelected"`
	HeapStyle   struct {
		StyleType      int `json:"styleType"`
		PromotionLabel struct {
			Text string `json:"text"`
		} `json:"promotionLabel"`
		Desc string `json:"desc"`
	} `json:"heapStyle"`
	GiftShowTop   bool   `json:"giftShowTop"`
	GiftCanDelete bool   `json:"giftCanDelete"`
	HeapType      string `json:"heapType"`
}

type CatDeliveryInfo struct {
	DeliveryPriceThreshold int    `json:"deliveryPriceThreshold"`
	ShippingFee            int    `json:"shippingFee"`
	ShippingFeeStr         string `json:"shippingFeeStr"`
	FreeThreshold          int    `json:"freeThreshold"`
	PinkageShow            bool   `json:"pinkageShow"`
	DeliveryStyle          struct {
		Desc string `json:"desc"`
	} `json:"deliveryStyle"`
}
type CartRefreshData struct {
	CartItems        []CartItem      `json:"cartItems"`
	ActivityItems    []ActivityItem  `json:"activityItems"`
	CatDeliveryInfo  CatDeliveryInfo `json:"catDeliveryInfo"`
	TotalAmount      int             `json:"totalAmount"`
	RealTotalAmount  int             `json:"realTotalAmount"`
	ReducedAmount    int             `json:"reducedAmount"`
	ReducedAmountV2  int             `json:"reducedAmountV2"`
	TotalDescription string          `json:"totalDescription"`
	TotalItemCounts  int             `json:"totalItemCounts"`
	AllItemsCounts   int             `json:"allItemsCounts"`
	CartDetailInfo   CartDetailInfo  `json:"cartDetailInfo"`
	AllSelected      bool            `json:"allSelected"`
	AddSuccess       bool            `json:"addSuccess"`
	OpResult         int             `json:"opResult"`
	PoiID            int             `json:"poiId"`
	StockPois        []int           `json:"stockPois"`
	NormalSkuList    []int           `json:"normalSkuList"`
	TagInfoMap       struct {
	} `json:"tagInfoMap"`
	SkuQtyList          []SkuQtyItem `json:"skuQtyList"`
	CacheID             string       `json:"cacheId"`
	CartReqID           string       `json:"cartReqId"`
	DisableSettleButton int          `json:"disableSettleButton"`
}

type CartDetailInfo struct {
	DiscountDesc    string `json:"discountDesc"`
	TotalAmount     int    `json:"totalAmount"`
	DiscountDetails []struct {
		TotalAmount     string `json:"totalAmount"`
		DiscountName    string `json:"discountName"`
		Type            int    `json:"type"`
		TotalAmountLine string `json:"totalAmountLine,omitempty"`
		DiscountTagText string `json:"discountTagText,omitempty"`
	} `json:"discountDetails"`
}

type SkuQtyItem struct {
	SkuID      int    `json:"skuId"`
	Quantity   string `json:"quantity"`
	QuantityV2 int    `json:"quantityV2"`
}

type ReponseError struct {
	code    int
	message string
}

func (r ReponseError) Error() string {
	return fmt.Sprintf("code=[%d] msg=[%s]", r.code, r.message)
}

// 叮咚

type DingdongResponse struct {
	Success bool            `json:"success"`
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

type CartData struct {
	NewOrderProductList []NewOrderProductItem `json:"new_order_product_list"`
	ParentOrderInfo     struct {
		ParentOrderSign string `json:"parent_order_sign"`
	} `json:"parent_order_info"`
}

type NewOrderProductItem struct {
	CanUsedBalanceMoney    string        `json:"can_used_balance_money"`
	CanUsedPointMoney      string        `json:"can_used_point_money"`
	CanUsedPointNum        int64         `json:"can_used_point_num"`
	CartCount              int64         `json:"cart_count"`
	FreeOrderCouponID      []interface{} `json:"freeOrderCouponId"`
	FrontPackageBgColor    string        `json:"front_package_bg_color"`
	FrontPackageStockColor string        `json:"front_package_stock_color"`
	FrontPackageText       string        `json:"front_package_text"`
	FrontPackageType       int64         `json:"front_package_type"`
	GoodsRealMoney         string        `json:"goods_real_money"`
	InstantRebateMoney     string        `json:"instant_rebate_money"`
	IsPresale              int64         `json:"is_presale"`
	IsShareStation         int64         `json:"is_share_station"`
	IsSupplyOrder          bool          `json:"is_supply_order"`
	OnlyTodayProducts      []interface{} `json:"only_today_products"`
	OnlyTomorrowProducts   []interface{} `json:"only_tomorrow_products"`
	PackageID              int64         `json:"package_id"`
	PackageType            int64         `json:"package_type"`
	Products               []struct {
		AccessoryInfoList        []interface{} `json:"accessoryInfoList"`
		ActivityDiscountInfoList []struct {
			ActivityID              string      `json:"activityId"`
			ActivityProductLimitNum interface{} `json:"activityProductLimitNum"`
			ActivityType            interface{} `json:"activityType"`
			CostType                interface{} `json:"costType"`
			Num                     int64       `json:"num"`
			Price                   string      `json:"price"`
			PriceType               int64       `json:"priceType"`
			ThresholdMoney          interface{} `json:"thresholdMoney"`
		} `json:"activityDiscountInfoList"`
		ActivityID                string `json:"activity_id"`
		BuyLimit                  int64  `json:"buy_limit"`
		CartID                    string `json:"cart_id"`
		CategoryPath              string `json:"category_path"`
		ConditionsNum             string `json:"conditions_num"`
		Count                     int64  `json:"count"`
		Description               string `json:"description"`
		ID                        string `json:"id"`
		InstantRebateMoney        string `json:"instant_rebate_money"`
		IsBooking                 int64  `json:"is_booking"`
		IsBulk                    int64  `json:"is_bulk"`
		IsGift                    int64  `json:"is_gift"`
		IsInvoice                 int64  `json:"is_invoice"`
		IsPresale                 int64  `json:"is_presale"`
		IsSharedStationProduct    int64  `json:"is_shared_station_product"`
		ManageCategoryPath        string `json:"manage_category_path"`
		NetWeight                 string `json:"net_weight"`
		NetWeightUnit             string `json:"net_weight_unit"`
		NoSupplementaryPrice      string `json:"no_supplementary_price"`
		NoSupplementaryTotalPrice string `json:"no_supplementary_total_price"`
		OrderSort                 int64  `json:"order_sort"`
		OriginPrice               string `json:"origin_price"`
		ParentBatchType           int64  `json:"parent_batch_type"`
		ParentID                  string `json:"parent_id"`
		Price                     string `json:"price"`
		PriceType                 int64  `json:"price_type"`
		ProductName               string `json:"product_name"`
		ProductType               int64  `json:"product_type"`
		PromotionNum              int64  `json:"promotion_num"`
		SaleBatches               struct {
			BatchType int64 `json:"batch_type"`
		} `json:"sale_batches"`
		SignType          int64         `json:"signType"`
		SizePrice         string        `json:"size_price"`
		Sizes             []interface{} `json:"sizes"`
		SkuActivityID     string        `json:"sku_activity_id"`
		SmallImage        string        `json:"small_image"`
		StorageValueID    int64         `json:"storage_value_id"`
		SubList           []interface{} `json:"sub_list"`
		SupplementaryList []interface{} `json:"supplementary_list"`
		TemperatureLayer  string        `json:"temperature_layer"`
		TotalOriginPrice  string        `json:"total_origin_price"`
		TotalPrice        string        `json:"total_price"`
		TotalMoney        string        `json:"total_money"`
		TotalOriginMoney  string        `json:"total_origin_money"`
		Type              int64         `json:"type"`
		ViewTotalWeight   string        `json:"view_total_weight"`
	} `json:"products"`
	StationID            string `json:"stationId"`
	TotalCount           int64  `json:"total_count"`
	TotalMoney           string `json:"total_money"`
	TotalOriginMoney     string `json:"total_origin_money"`
	TotalRebateMoney     string `json:"total_rebate_money"`
	UsedBalanceMoney     string `json:"used_balance_money"`
	UsedPointMoney       string `json:"used_point_money"`
	UsedPointNum         int64  `json:"used_point_num"`
	ReservedTimeStart    int64  `json:"reserved_time_start"`
	ReservedTimeEnd      int64  `json:"reserved_time_end"`
	SoonArrival          string `json:"soon_arrival"`
	FirstSelectedBigTime int    `json:"first_selected_big_time"`
}

type OrderData struct {
	AlipayTicket             interface{}   `json:"alipay_ticket"`
	AllCouponCount           int64         `json:"all_coupon_count"`
	AllGoodsSavedMoney       string        `json:"all_goods_saved_money"`
	AllGoodsSavedMoneyNote   string        `json:"all_goods_saved_money_note"`
	CanUseFreightCoupon      bool          `json:"can_use_freight_coupon"`
	CanUseFreightCouponCount int64         `json:"can_use_freight_coupon_count"`
	CanUsedBalanceMoney      string        `json:"can_used_balance_money"`
	CanUsedPointMoney        string        `json:"can_used_point_money"`
	CanUsedPointNum          int64         `json:"can_used_point_num"`
	Coupons                  interface{}   `json:"coupons"`
	CouponsMoney             string        `json:"coupons_money"`
	DefaultCoupon            struct{}      `json:"default_coupon"`
	DefaultCoupons           []interface{} `json:"default_coupons"`
	DefaultFreightCoupon     struct{}      `json:"default_freight_coupon"`
	DefaultFreightCoupons    []interface{} `json:"default_freight_coupons"`
	DisableCouponImg         string        `json:"disable_coupon_img"`
	DisableCouponNote        string        `json:"disable_coupon_note"`
	DiscountMoney            string        `json:"discount_money"`
	DisplayTotalMoney        string        `json:"display_total_money"`
	FreeFreightNotice        string        `json:"free_freight_notice"`
	FreightDiscountMoney     string        `json:"freight_discount_money"`
	FreightMoney             string        `json:"freight_money"`
	FreightRealMoney         string        `json:"freight_real_money"`
	Freights                 []struct {
		Freight struct {
			DiscountFreightMoney string `json:"discount_freight_money"`
			FreightMoney         string `json:"freight_money"`
			FreightRealMoney     string `json:"freight_real_money"`
			Remark               string `json:"remark"`
			Type                 int64  `json:"type"`
		} `json:"freight"`
		PackageID            int64 `json:"package_id"`
		RealMatchSupplyOrder bool  `json:"real_match_supply_order"`
	} `json:"freights"`
	FullToOff              string      `json:"full_to_off"`
	GoodsOriginMoney       string      `json:"goods_origin_money"`
	GoodsRealMoney         string      `json:"goods_real_money"`
	InstantRebateMoney     string      `json:"instant_rebate_money"`
	InvoiceMoney           string      `json:"invoice_money"`
	PointHint              string      `json:"point_hint"`
	ProductShortOrderText  interface{} `json:"product_short_order_text"`
	ShowPoint              int64       `json:"show_point"`
	TotalMoney             string      `json:"total_money"`
	UsableCouponCount      int64       `json:"usable_coupon_count"`
	UseFreightCouponNotice interface{} `json:"use_freight_coupon_notice"`
	UsedBalanceMoney       string      `json:"used_balance_money"`
	UsedPointMoney         string      `json:"used_point_money"`
	UsedPointNum           int64       `json:"used_point_num"`
	UserPointNum           int64       `json:"user_point_num"`
	Vip                    struct{}    `json:"vip"`
	VipFreeFreight         interface{} `json:"vip_free_freight"`
	VipGoodsSaveMoney      string      `json:"vip_goods_save_money"`
	VipGoodsSaveMoneyNote  interface{} `json:"vip_goods_save_money_note"`
	VipMoney               string      `json:"vip_money"`
}
type checkOrderData struct {
	Order OrderData `json:"order"`
}

type getMultiReserveTimeItem struct {
	AllFull bool `json:"all_full"`
	Time    []struct {
		DateStr          string            `json:"date_str"`
		DateStrTimestamp int64             `json:"date_str_timestamp"`
		Day              string            `json:"day"`
		IsInvalid        bool              `json:"is_invalid"`
		TimeFullTextTip  string            `json:"time_full_text_tip"`
		Times            []ReserveTimeItem `json:"times"`
	} `json:"time"`
}

type ReserveTimeItem struct {
	ArrivalTime             bool        `json:"arrival_time"`
	ArrivalTimeMsg          string      `json:"arrival_time_msg"`
	DisableMsg              string      `json:"disableMsg"`
	DisableType             int64       `json:"disableType"`
	EndTime                 string      `json:"end_time"`
	EndTimestamp            int64       `json:"end_timestamp"`
	FullFlag                bool        `json:"fullFlag"`
	SelectMsg               string      `json:"select_msg"`
	StartTime               string      `json:"start_time"`
	StartTimestamp          int64       `json:"start_timestamp"`
	SupplyOrderCountDownTip interface{} `json:"supply_order_count_down_tip"`
	SupplyOrderEndTime      interface{} `json:"supply_order_end_time"`
	SupplyOrderTip          interface{} `json:"supply_order_tip"`
	TextMsg                 string      `json:"textMsg"`
	TimeBizType             interface{} `json:"time_biz_type"`
	Type                    int64       `json:"type"`
}
