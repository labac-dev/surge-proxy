package reddit

import (
	"encoding/json"
	"net/http"

	"github.com/labac-dev/surge-proxy/common"
)

func ModifyResponse(res *http.Response) error {
	var (
		data any
		err  error
		body []byte
	)

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	cleanedData := RemoveAds(data)

	body, err = json.Marshal(cleanedData)
	if err != nil {
		return err
	}

	return common.Response(res, body)
}

func IsAds(data map[string]any) bool {
	if data["adPayload"] != nil || data["adEvents"] != nil {
		return true
	}

	if v, ok := data["__typename"].(string); ok && v == "AdPost" {
		return true
	}

	return false
}

func DisableGilding(data map[string]any) {
	if isGildable, ok := data["isGildable"].(bool); ok && isGildable {
		data["isGildable"] = false
	}
}

func RemoveAds(data any) any {
	switch v := data.(type) {
	case []any:
		var filtered = make([]any, 0, len(v))
		for _, item := range v {
			if result := RemoveAds(item); result != nil {
				filtered = append(filtered, result)
			}
		}
		return filtered
	case map[string]any:
		// Check if the node is an ad
		if node, ok := v["node"].(map[string]any); ok && IsAds(node) {
			return nil
		}

		// // Check if the data is an ad
		if IsAds(v) {
			return nil
		}

		// Remove the Gilable button
		DisableGilding(v)

		for key, value := range v {
			v[key] = RemoveAds(value)
		}
	}

	return data
}
