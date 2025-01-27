package busuu

import (
	"encoding/json"
	"net/http"

	"github.com/labac-dev/surge-proxy/internal/common"
)

func ModifyResponse(res *http.Response) error {
	if res.Request.URL.Path == "/users/me" {
		var data interface{}

		err := json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			return err
		}

		dataMap := data.(map[string]interface{})
		dataMap["data"].(map[string]interface{})["premium_data"] = map[string]interface{}{
			"product":          nil,
			"created":          1737324382,
			"started":          nil,
			"next_charge":      nil,
			"expiration":       1893456000,
			"cancelled":        nil,
			"platform":         nil,
			"productId":        "com.busuu.app.subs12month.serious_ft_jan_24_variant_a",
			"market":           "AppStore",
			"subscriptionType": "1y",
			"type":             "mobile",
			"subscriptionId":   nil,
			"providerId":       nil,
			"current_plan": map[string]interface{}{
				"id":             "com.busuu.app.subs12month.serious_ft_jan_24_variant_a",
				"name":           12,
				"months":         12,
				"saving":         0,
				"active":         true,
				"billing_period": "month",
				"tier":           "plus",
				"price":          map[string]interface{}{"amount": nil, "price_human": nil, "currencyCode": ""},
				"monthlyPrice":   map[string]interface{}{"amount": nil, "price_human": nil, "currencyCode": ""},
			},
			"upgrade_plan":        nil,
			"suspended_plan":      nil,
			"partner":             nil,
			"partner_metadata":    nil,
			"in_app_cancellation": false,
			"subscription_status": "active",
			"free_trial":          map[string]interface{}{"free_trial": false},
		}
		dataMap["data"].(map[string]interface{})["is_premium"] = true
		dataMap["data"].(map[string]interface{})["roles"] = []int{6, 32}
		dataMap["data"].(map[string]interface{})["free_trial_eligible"] = false
		dataMap["data"].(map[string]interface{})["access"] = map[string]string{"tier": "plus"}

		body, err := json.Marshal(data)
		if err != nil {
			return err
		}

		return common.Response(res, body)
	}

	return nil
}
