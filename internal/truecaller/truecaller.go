package truecaller

import (
	"encoding/json"
	"net/http"

	"github.com/labac-dev/surge-proxy/internal/common"
)

type Product struct {
	ID          string `json:"id"`
	SKU         string `json:"sku"`
	ContentType string `json:"contentType"`
	ProductType string `json:"productType"`
	IsFreeTrial bool   `json:"isFreeTrial"`
}

type Feature struct {
	ID     string `json:"id"`
	Rank   int    `json:"rank"`
	Status string `json:"status"`
	IsFree bool   `json:"isFree"`
}

type Tier struct {
	ID      string    `json:"id"`
	Feature []Feature `json:"feature"`
}

type StatusResponse struct {
	Expire               string  `json:"expire"`
	Start                string  `json:"start"`
	PaymentProvider      string  `json:"paymentProvider"`
	IsExpired            bool    `json:"isExpired"`
	IsGracePeriodExpired bool    `json:"isGracePeriodExpired"`
	InAppPurchaseAllowed bool    `json:"inAppPurchaseAllowed"`
	Product              Product `json:"product"`
	Tier                 Tier    `json:"tier"`
	Scope                string  `json:"scope"`
}

type PlanResponse struct {
	Tier []interface{} `json:"tier"`
}

var statusResponse = StatusResponse{
	Expire:               "2030-01-01T00:00:00Z",
	Start:                "2025-01-01T00:00:00Z",
	PaymentProvider:      "Apple",
	IsExpired:            false,
	IsGracePeriodExpired: false,
	InAppPurchaseAllowed: false,
	Product: Product{
		ID:          "renewable.premiumgold.annual",
		SKU:         "renewable.premiumgold.annual",
		ContentType: "subscription",
		ProductType: "GoldYearly",
		IsFreeTrial: false,
	},
	Tier: Tier{
		ID: "gold",
		Feature: []Feature{
			{ID: "premium_feature", Rank: -2147483648, Status: "Included", IsFree: false},
			{ID: "no_ads", Rank: 1, Status: "Included", IsFree: false},
			{ID: "extended_spam_blocking", Rank: 2, Status: "Included", IsFree: false},
			{ID: "advanced_caller_id", Rank: 3, Status: "Included", IsFree: false},
			{ID: "live_lookup", Rank: 4, Status: "Included", IsFree: false},
			{ID: "verified_badge", Rank: 5, Status: "Included", IsFree: false},
			{ID: "spam_stats", Rank: 6, Status: "Included", IsFree: false},
			{ID: "auto_spam_block", Rank: 7, Status: "Included", IsFree: false},
			{ID: "call_alert", Rank: 8, Status: "Included", IsFree: false},
			{ID: "identifai", Rank: 14, Status: "Included", IsFree: false},
			{ID: "siri_search", Rank: 15, Status: "Included", IsFree: false},
			{ID: "who_viewed_my_profile", Rank: 16, Status: "Included", IsFree: false},
			{ID: "incognito_mode", Rank: 19, Status: "Included", IsFree: false},
			{ID: "premium_badge", Rank: 20, Status: "Included", IsFree: false},
			{ID: "premium_app_icon", Rank: 21, Status: "Included", IsFree: false},
			{ID: "live_chat_support", Rank: 23, Status: "Included", IsFree: false},
			{ID: "premium_support", Rank: 24, Status: "Included", IsFree: false},
			{ID: "family_sharing", Rank: 25, Status: "Excluded", IsFree: false},
			{ID: "gold_caller_id", Rank: 26, Status: "Included", IsFree: false},
		},
	},
	Scope: "paid_gold",
}

var planResponse = PlanResponse{
	Tier: []interface{}{},
}

func ModifyResponse(res *http.Response) error {
	if res.Request.URL.Path == "/v3/subscriptions/status" {
		responseBytes, err := json.Marshal(statusResponse)
		if err != nil {
			return err
		}
		return common.Response(res, responseBytes)
	}

	if res.Request.URL.Path == "/v7/products/apple" {
		responseBytes, err := json.Marshal(planResponse)
		if err != nil {
			return err
		}
		return common.Response(res, responseBytes)
	}

	return nil
}
