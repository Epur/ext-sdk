//https://shopify.dev/docs/apps/auth/get-access-tokens/token-exchange/getting-started#step-3-make-authenticated-requests
/** 在线交换令牌
curl -X POST \
https://{shop}.myshopify.com/admin/oauth/access_token \
-H 'Content-Type: application/json' \
-H 'Accept: application/json' \
-d '{
"client_id": "{client_id}",
"client_secret": "{client_secret}",
"grant_type": "urn:ietf:params:oauth:grant-type:token-exchange",
"subject_token": "{session_token}",
"subject_token_type": "urn:ietf:params:oauth:token-type:id_token",
"requested_token_type": "urn:shopify:params:oauth:token-type:online-access-token"
}'
离线交换令牌
curl -X POST \
https://{shop}.myshopify.com/admin/oauth/access_token \
-H 'Content-Type: application/json' \
-H 'Accept: application/json' \
-d '{
"client_id": "{client_id}",
"client_secret": "{client_secret}",
"grant_type": "urn:ietf:params:oauth:grant-type:token-exchange",
"subject_token": "{session_token}",
"subject_token_type": "urn:ietf:params:oauth:token-type:id_token",
"requested_token_type": "urn:shopify:params:oauth:token-type:offline-access-token"
}'

**/
//离线访问令牌
https://shopify.dev/docs/apps/build/authentication-authorization/access-token-types/offline-access-tokens

//订单
https://shopify.dev/docs/api/admin-rest/2024-04/resources/order#get-orders?status=any
HTTP/1.1 200 OK
{
"orders": [
{
"id": 450789469,
"created_at": "2008-01-10T11:00:00-05:00",
"name": "#1001",
"total_price": "598.94"
}
]
}
//订单详情
https://shopify.dev/docs/api/admin-rest/2024-04/resources/order#get-orders-order-id?fields=id,line-items,name,total-price
HTTP/1.1 200 OK
{
"order": {
"id": 450789469,
"name": "#1001",
"total_price": "598.94",
"line_items": [
{
"id": 466157049,
"admin_graphql_api_id": "gid://shopify/LineItem/466157049",
"current_quantity": 0,
"fulfillable_quantity": 0,
"fulfillment_service": "manual",
"fulfillment_status": null,
"gift_card": false,
"grams": 200,
"name": "IPod Nano - 8gb - green",
"price": "199.00",
"price_set": {
"shop_money": {
"amount": "199.00",
"currency_code": "USD"
},
"presentment_money": {
"amount": "199.00",
"currency_code": "USD"
}
},
"product_exists": true,
"product_id": 632910392,
"properties": [
{
"name": "Custom Engraving Front",
"value": "Happy Birthday"
},
{
"name": "Custom Engraving Back",
"value": "Merry Christmas"
}
],
"quantity": 1,
"requires_shipping": true,
"sku": "IPOD2008GREEN",
"taxable": true,
"title": "IPod Nano - 8gb",
"total_discount": "0.00",
"total_discount_set": {
"shop_money": {
"amount": "0.00",
"currency_code": "USD"
},
"presentment_money": {
"amount": "0.00",
"currency_code": "USD"
}
},
"variant_id": 39072856,
"variant_inventory_management": "shopify",
"variant_title": "green",
"vendor": null,
"tax_lines": [
{
"channel_liable": null,
"price": "3.98",
"price_set": {
"shop_money": {
"amount": "3.98",
"currency_code": "USD"
},
"presentment_money": {
"amount": "3.98",
"currency_code": "USD"
}
},
"rate": 0.06,
"title": "State Tax"
}
],
"duties": [],
"discount_allocations": [
{
"amount": "3.34",
"amount_set": {
"shop_money": {
"amount": "3.34",
"currency_code": "USD"
},
"presentment_money": {
"amount": "3.34",
"currency_code": "USD"
}
},
"discount_application_index": 0
}
]
},
{
"id": 518995019,
"admin_graphql_api_id": "gid://shopify/LineItem/518995019",
"current_quantity": 1,
"fulfillable_quantity": 1,
"fulfillment_service": "manual",
"fulfillment_status": null,
"gift_card": false,
"grams": 200,
"name": "IPod Nano - 8gb - red",
"price": "199.00",
"price_set": {
"shop_money": {
"amount": "199.00",
"currency_code": "USD"
},
"presentment_money": {
"amount": "199.00",
"currency_code": "USD"
}
},
"product_exists": true,
"product_id": 632910392,
"properties": [],
"quantity": 1,
"requires_shipping": true,
"sku": "IPOD2008RED",
"taxable": true,
"title": "IPod Nano - 8gb",
"total_discount": "0.00",
"total_discount_set": {
"shop_money": {
"amount": "0.00",
"currency_code": "USD"
},
"presentment_money": {
"amount": "0.00",
"currency_code": "USD"
}
},
"variant_id": 49148385,
"variant_inventory_management": "shopify",
"variant_title": "red",
"vendor": null,
"tax_lines": [
{
"channel_liable": null,
"price": "3.98",
"price_set": {
"shop_money": {
"amount": "3.98",
"currency_code": "USD"
},
"presentment_money": {
"amount": "3.98",
"currency_code": "USD"
}
},
"rate": 0.06,
"title": "State Tax"
}
],
"duties": [],
"discount_allocations": [
{
"amount": "3.33",
"amount_set": {
"shop_money": {
"amount": "3.33",
"currency_code": "USD"
},
"presentment_money": {
"amount": "3.33",
"currency_code": "USD"
}
},
"discount_application_index": 0
}
]
},
{
"id": 703073504,
"admin_graphql_api_id": "gid://shopify/LineItem/703073504",
"current_quantity": 0,
"fulfillable_quantity": 0,
"fulfillment_service": "manual",
"fulfillment_status": null,
"gift_card": false,
"grams": 200,
"name": "IPod Nano - 8gb - black",
"price": "199.00",
"price_set": {
"shop_money": {
"amount": "199.00",
"currency_code": "USD"
},
"presentment_money": {
"amount": "199.00",
"currency_code": "USD"
}
},
"product_exists": true,
"product_id": 632910392,
"properties": [],
"quantity": 1,
"requires_shipping": true,
"sku": "IPOD2008BLACK",
"taxable": true,
"title": "IPod Nano - 8gb",
"total_discount": "0.00",
"total_discount_set": {
"shop_money": {
"amount": "0.00",
"currency_code": "USD"
},
"presentment_money": {
"amount": "0.00",
"currency_code": "USD"
}
},
"variant_id": 457924702,
"variant_inventory_management": "shopify",
"variant_title": "black",
"vendor": null,
"tax_lines": [
{
"channel_liable": null,
"price": "3.98",
"price_set": {
"shop_money": {
"amount": "3.98",
"currency_code": "USD"
},
"presentment_money": {
"amount": "3.98",
"currency_code": "USD"
}
},
"rate": 0.06,
"title": "State Tax"
}
],
"duties": [],
"discount_allocations": [
{
"amount": "3.33",
"amount_set": {
"shop_money": {
"amount": "3.33",
"currency_code": "USD"
},
"presentment_money": {
"amount": "3.33",
"currency_code": "USD"
}
},
"discount_application_index": 0
}
]
}
]
}
}


//检索店铺
HTTP/1.1 200 OK
{
"shop": {
"id": 548380009,
"name": "John Smith Test Store",
"email": "j.smith@example.com",
"domain": "shop.apple.com",
"province": "California",
"country": "US",
"address1": "1 Infinite Loop",
"zip": "95014",
"city": "Cupertino",
"source": null,
"phone": "1231231234",
"latitude": 45.45,
"longitude": -75.43,
"primary_locale": "en",
"address2": "Suite 100",
"created_at": "2007-12-31T19:00:00-05:00",
"updated_at": "2024-04-01T13:38:18-04:00",
"country_code": "US",
"country_name": "United States",
"currency": "USD",
"customer_email": "customers@apple.com",
"timezone": "(GMT-05:00) Eastern Time (US & Canada)",
"iana_timezone": "America/New_York",
"shop_owner": "John Smith",
"money_format": "${{amount}}",
"money_with_currency_format": "${{amount}} USD",
"weight_unit": "lb",
"province_code": "CA",
"taxes_included": null,
"auto_configure_tax_inclusivity": null,
"tax_shipping": null,
"county_taxes": true,
"plan_display_name": "Shopify Plus",
"plan_name": "enterprise",
"has_discounts": true,
"has_gift_cards": true,
"myshopify_domain": "jsmith.myshopify.com",
"google_apps_domain": null,
"google_apps_login_enabled": null,
"money_in_emails_format": "${{amount}}",
"money_with_currency_in_emails_format": "${{amount}} USD",
"eligible_for_payments": true,
"requires_extra_payments_agreement": false,
"password_enabled": false,
"has_storefront": true,
"finances": true,
"primary_location_id": 655441491,
"checkout_api_supported": true,
"multi_location_enabled": true,
"setup_required": false,
"pre_launch_enabled": false,
"enabled_presentment_currencies": [
"USD"
],
"transactional_sms_disabled": false,
"marketing_sms_consent_enabled_at_checkout": false
}
}