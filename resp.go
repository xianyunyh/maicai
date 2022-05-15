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
