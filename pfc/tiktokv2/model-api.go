package tiktokv2

type GetTokenResponse struct {
	AccessToken          string `json:"access_token"`            //访问令牌
	RefreshToken         string `json:"refresh_token"`           //刷新令牌
	AccessTokenExpireIn  int64  `json:"access_token_expire_in"`  //访问令牌有效期
	RefreshTokenExpireIn int64  `json:"refresh_token_expire_in"` //刷新令牌有效期
	OpenId               string `json:"open_id"`                 //用户唯一标识
	SellerName           string `json:"seller_name"`             //卖家名称
	SellerBaseRegion     string `json:"seller_base_region"`      //卖家所在区域，如ID:印尼
	UserType             int    `json:"user_type"`               //卖家所在区域，如ID:印尼
}

// 获取店铺信息
type ShopListResponse struct {
	Shop []Shop `json:"shops"`
}
type Shop struct {
	ShopCipher string `json:"cipher"`      // 商店密码(Tk用)
	Code       string `json:"code"`        // 店铺code
	ShopId     string `json:"id"`          // 店铺ID
	ShopName   string `json:"name"`        // 店铺名称
	Region     string `json:"region"`      // 店铺所在区域
	SellerType string `json:"seller_type"` // 跨境类型: LOCAL-本土店 CROSS_BORDER-跨境店
}

// 返回与卖家关联的所有商店
type ActiveShopListResponse struct {
	Shop []ActiveShop `json:"shops"`
}
type ActiveShop struct {
	Id     string `json:"id"`
	Region string `json:"region"`
}

type GetOrderDetailResponse struct {
	List []OrderDetailResponse `json:"order_list"`
}

type GetOrderListResponse struct {
	Total int            `json:"total"`
	List  []OrderListRow `json:"list"`
}

type GetProductListResponse struct {
	NextPageToken string `json:"next_page_token"`
	Products      []struct {
		CreateTime             int      `json:"create_time"`
		Id                     string   `json:"id"`
		ProductSyncFailReasons string   `json:"product_sync_fail_reasons"`
		SalesRegions           []string `json:"sales_regions"`
		Skus                   []struct {
			Id        string `json:"id"`
			Inventory []struct {
				Quantity    int    `json:"quantity"`
				WarehouseId string `json:"warehouse_id"`
			} `json:"inventory"`
			Price struct {
				Currency          string `json:"currency"`
				SalePrice         string `json:"sale_price"`
				TaxExclusivePrice string `json:"tax_exclusive_price"`
			} `json:"price"`
			SellerSku string `json:"seller_sku"`
		} `json:"skus"`
		Status     string `json:"status"`
		Title      string `json:"title"`
		UpdateTime int    `json:"update_time"`
	} `json:"products"`
	TotalCount int `json:"total_count"`
}

type GetProductDetailResponse struct {
	ProductDetailResponse
}

type GetOrderPrintAwbResponse struct {
	DocUrl string `json:"doc_url"`
}

type GetGlobalProductDetailResponse struct {
	GlobalProductDetailResponse
}

type CreateProductResponse struct {
	CreateProduct
}

type CreateGlobalProductResponse struct {
	GlobalProductId string      `json:"global_product_id"`
	GlobalSkus      []GlobalSku `json:"global_skus"`
}

type UploadProductImageResponse struct {
	UploadProductImage
}

type PublishGlobalProductResponse struct {
	PublishGlobalProduct
}

type CategoriesAttributesResponse struct {
	CategoriesAttributes
}

type GetGlobalProductsCategoriesResponse struct {
	Categories []GetGlobalProductsCategories `json:"categories"`
}

type GetGlobalProductListResponse struct {
	GlobalProducts []GlobalProduct `json:"global_products"`
	NextPageToken  string          `json:"next_page_token"`
	TotalCount     int             `json:"total_count"`
}

type EditGlobalProductResponse struct {
	GlobalSkus []struct {
		Id              string `json:"id"`
		SalesAttributes []struct {
			Id      string `json:"id"`
			ValueId string `json:"value_id"`
		} `json:"sales_attributes"`
		SellerSku string `json:"seller_sku"`
	} `json:"global_skus"`
	PublishResults []struct {
		FailReasons []struct {
			Message string `json:"message"`
		} `json:"fail_reasons"`
		Region string `json:"region"`
		Status string `json:"status"`
	} `json:"publish_results"`
}

type ActivateProductResponse struct {
	Errors []struct {
		Code   int `json:"code"`
		Detail struct {
			ExtraErrors []struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"extra_errors"`
			ProductId string `json:"product_id"`
		} `json:"detail"`
		Message string `json:"message"`
	} `json:"errors"`
}

type DeactivateProductResponse struct {
	Errors []struct {
		Code   int `json:"code"`
		Detail struct {
			ProductId string `json:"product_id"`
		} `json:"detail"`
		Message string `json:"message"`
	} `json:"errors"`
}
