package constant

const (
	/**token-exchange (交换令牌)**/
	SERVER_URl = "https://%s.myshopify.com"
	AUTHSITE   = "https://%s.myshopify.com"

	EXCHGACCESS  = "/admin/oauth/access_token"
	GETACCESS    = "/admin/oauth/access_token"
	REFRESHTOKEN = ""
	//{shop}	用户的商店名称。
	AUTH = "/admin/oauth/authorize"

	//检索订单列表
	ORDER_LIST = "/admin/api/%s/orders.json?status=any"

	//获取订单明细
	ORDER_DETAIL = "/admin/api/%s/orders/%s.json"

	//检索店铺
	SHOP_CONFIG = "/admin/api/%s/shop.json"
)
