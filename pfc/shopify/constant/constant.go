package constant

const (
	AUTH_SCOPE                 = "write_products,read_shipping,write_orders"
	SUBJECT_TOKEN_TYPE_DEFAULT = "urn:ietf:params:oauth:token-type:id_token"
	REQUEST_TOKEN_TYPE_ON      = "urn:shopify:params:oauth:token-type:online-access-token"
	REQUEST_TOKEN_TOKEN_OFF    = "urn:shopify:params:oauth:token-type:offline-access-token"

	//检索订单字段集
	ORDER_LIST_FIELDS = "created_at,id,name,total-price"

	//检索订单详情字段集
	ORDER_DETAIL_FIELDS = "id,line_items,name,total_price"

	//店铺检索字段集
	SHOP_FIELDS = "address1,address2,city,province,country"
)
const (
	API_V1 = "2024-04"
)
