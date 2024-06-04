package main

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/pfc/shein"
)

func main() {
	orderDetail := `{
    "code": "0",
    "msg": "OK",
    "info": [
        {
            "orderNo": "GSONK315800M5DC",
            "orderType": 1,
            "performanceType": 1,
            "orderStatus": 2,
            "isCod": 2,
            "isOverLimitOrder": 1,
            "unpackingStatus": 3,
            "orderTag": 0,
            "deliveryType": 1,
            "printOrderStatus": 1,
            "invoiceStatus": 2,
            "orderGoodsInfoList": [
                {
                    "goodsId": 3000000000232252,
                    "skuCode": "I5dtvb4nk1dx",
                    "skc": "sM23040300586693",
                    "goodsSn": "SPMP发布自定义保存草稿2",
                    "sellerSku": "",
                    "goodsStatus": 1,
                    "newGoodsStatus": 2,
                    "skuAttribute": [
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "CN"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "US"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7葡文-次规格值自定义2",
                            "language": "PT"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7泰语-次规格值自定义2",
                            "language": "TH"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "ES"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "IT"
                        }
                    ],
                    "goodsTitle": "GoodsName111111111111",
                    "spuPicURL": "http://imgdeal-test01.shein.com/pi_img/2021/11/03/16359230894287290938_thumbnail_220x293.jpg",
                    "goodsWeight": 500.00,
                    "storageTag": 1,
                    "performanceTag": 2,
                    "goodsExchangeTag": 1,
                    "unpackingGroupNo": "FG24022932245297156",
                    "unpackingGroupInvoiceStatus": 2,
                    "beExchangeEntityId": 0,
                    "orderCurrency": "MXN",
                    "sellerCurrencyPrice": 280.45,
                    "orderCurrencyStoreCouponPrice": 0.00,
                    "orderCurrencyPromotionPrice": 0.00,
                    "commission": 0.00,
                    "commissionRate": 0.0000,
                    "serviceCharge": 0.00,
                    "performanceServiceCharge": 0.00,
                    "estimatedIncome": 280.45,
                    "spuName": "M2304030058",
                    "saleTax": 0.00,
                    "warehouseCode": null,
                    "warehouseName": null,
                    "sellerCurrencyDiscountPrice": 801.06
                },
                {
                    "goodsId": 3000000000232253,
                    "skuCode": "I5dtvb4nk1dx",
                    "skc": "sM23040300586693",
                    "goodsSn": "SPMP发布自定义保存草稿2",
                    "sellerSku": "",
                    "goodsStatus": 1,
                    "newGoodsStatus": 2,
                    "skuAttribute": [
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "CN"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "US"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7葡文-次规格值自定义2",
                            "language": "PT"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7泰语-次规格值自定义2",
                            "language": "TH"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "ES"
                        },
                        {
                            "attrValueId": "1007750,39,1007752",
                            "attrName": "主规格值自定义2-7-次规格值自定义2",
                            "language": "IT"
                        }
                    ],
                    "goodsTitle": "GoodsName111111111111",
                    "spuPicURL": "http://imgdeal-test01.shein.com/pi_img/2021/11/03/16359230894287290938_thumbnail_220x293.jpg",
                    "goodsWeight": 500.00,
                    "storageTag": 1,
                    "performanceTag": 2,
                    "goodsExchangeTag": 1,
                    "unpackingGroupNo": "FG24022932245297155",
                    "unpackingGroupInvoiceStatus": 2,
                    "beExchangeEntityId": 0,
                    "orderCurrency": "MXN",
                    "sellerCurrencyPrice": 280.45,
                    "orderCurrencyStoreCouponPrice": 0.00,
                    "orderCurrencyPromotionPrice": 0.00,
                    "commission": 0.00,
                    "commissionRate": 0.0000,
                    "serviceCharge": 0.00,
                    "performanceServiceCharge": 0.00,
                    "estimatedIncome": 280.45,
                    "spuName": "M2304030058",
                    "saleTax": 0.00,
                    "warehouseCode": null,
                    "warehouseName": null,
                    "sellerCurrencyDiscountPrice": 801.06
                }
            ],
            "packageWaybillList": [
                {
                    "packageNo": "GC24030727444473856",
                    "waybillNo": "UAT301009158929",
                    "carrier": "222222",
                    "carrierCode": "TYDHL-BA1",
                    "productInventoryList": [
                        {
                            "productId": "3000000000232253"
                        }
                    ],
                    "packageLabel": "",
                    "sortingCode": "H09",
                    "expressSortingCode": "NYN",
                    "isCutOffSeller": 2
                },
                {
                    "packageNo": "GC24030727444473857",
                    "waybillNo": "UAT301008583535",
                    "carrier": "222222",
                    "carrierCode": "TYDHL-BA1",
                    "productInventoryList": [
                        {
                            "productId": "3000000000232252"
                        }
                    ],
                    "packageLabel": "",
                    "sortingCode": "H09",
                    "expressSortingCode": "NYN",
                    "isCutOffSeller": 2
                }
            ],
            "orderCurrency": "MXN",
            "productTotalPrice": 560.90,
            "storeDiscountTotalPrice": 0.00,
            "promotionDiscountTotalPrice": 0.00,
            "totalCommission": 0.00,
            "totalServiceCharge": 0.00,
            "totalPerformanceServiceCharge": 0.00,
            "estimatedGrossIncome": 560.90,
            "totalSaleTax": 0.00,
            "orderTime": "2024-02-29T17:03:31.000+0800",
            "paymentTime": "2024-02-29T17:03:32.000+0800",
            "orderAllocateTime": "2024-02-29T17:05:05.000+0800",
            "requestDeliveryTime": "2024-03-01T17:05:05.000+0800",
            "sellerDeliveryTime": "",
            "warehouseDeliveryTime": "",
            "printingTime": "2024-03-07T17:00:55.000+0800",
            "scheduleDeliveryTime": "",
            "pickUpTime": "",
            "orderReceiptTime": "",
            "orderRejectionTime": "",
            "orderReportedLossTime": "",
            "orderReturnTime": "",
            "orderMsgUpdateTime": "2024-03-07T17:00:58.360+0800",
            "billNo": "GSONK315800M5DC",
            "salesArea": 1,
            "stockMode": 3,
            "salesSite": "shein-mx",
            "storeCode": 3342233928,
            "settleActuallyPrice": null,
            "unProcessReason": [],
            "packageInvoiceProblems": [
                {
                    "problemCode": null,
                    "problemDescEnglish": "The order package requires a CTE logistics invoice, and the platform logistics provider is currently invoicing. Please be patient and wait",
                    "problemField": null,
                    "proposalEnglish": null,
                    "packageNo": "GC24030727444473856"
                },
                {
                    "problemCode": null,
                    "problemDescEnglish": "The order package requires a CTE logistics invoice, and the platform logistics provider is currently invoicing. Please be patient and wait",
                    "problemField": null,
                    "proposalEnglish": null,
                    "packageNo": "GC24030727444473857"
                }
            ],
            "expectedCollectTime": ""
        }
    ],
    "bbl": {}
}`
	response := shein.OrderDetailResponse{}

	err := json.Unmarshal([]byte(orderDetail), &response)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%#v\n", response)
		//fmt.Println(response)
	}
	fmt.Println("===================")
}
