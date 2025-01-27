package truecaller

import (
	"net/http"

	"github.com/labac-dev/surge-proxy/internal/common"
)

const (
	STATUS_RESPONSE = `{"expire":"2030-01-01T00:00:00Z","start":"2025-01-01T00:00:00Z","paymentProvider":"Apple","isExpired":false,"isGracePeriodExpired":false,"inAppPurchaseAllowed":false,"product":{"id":"renewable.premiumgold.annual","sku":"renewable.premiumgold.annual","contentType":"subscription","productType":"GoldYearly","isFreeTrial":false},"tier":{"id":"gold","feature":[{"id":"premium_feature","rank":-2147483648,"status":"Included","isFree":false},{"id":"no_ads","rank":1,"status":"Included","isFree":false},{"id":"extended_spam_blocking","rank":2,"status":"Included","isFree":false},{"id":"advanced_caller_id","rank":3,"status":"Included","isFree":false},{"id":"live_lookup","rank":4,"status":"Included","isFree":false},{"id":"verified_badge","rank":5,"status":"Included","isFree":false},{"id":"spam_stats","rank":6,"status":"Included","isFree":false},{"id":"auto_spam_block","rank":7,"status":"Included","isFree":false},{"id":"call_alert","rank":8,"status":"Included","isFree":false},{"id":"identifai","rank":14,"status":"Included","isFree":false},{"id":"siri_search","rank":15,"status":"Included","isFree":false},{"id":"who_viewed_my_profile","rank":16,"status":"Included","isFree":false},{"id":"incognito_mode","rank":19,"status":"Included","isFree":false},{"id":"premium_badge","rank":20,"status":"Included","isFree":false},{"id":"premium_app_icon","rank":21,"status":"Included","isFree":false},{"id":"live_chat_support","rank":23,"status":"Included","isFree":false},{"id":"premium_support","rank":24,"status":"Included","isFree":false},{"id":"family_sharing","rank":25,"status":"Excluded","isFree":false},{"id":"gold_caller_id","rank":26,"status":"Included","isFree":false}]},"scope":"paid_gold"}`
	PLAN_RESPONSE   = `{"tier":[]}`
)

func ModifyResponse(res *http.Response) error {
	if res.Request.URL.Path == "/v3/subscriptions/status" {
		return common.Response(res, []byte(STATUS_RESPONSE))
	}

	if res.Request.URL.Path == "/v7/products/apple" {
		return common.Response(res, []byte(PLAN_RESPONSE))
	}

	return nil
}
