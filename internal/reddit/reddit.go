package reddit

import (
	"encoding/json"
	"net/http"

	"github.com/labac-dev/surge-proxy/internal/common"
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

	body, err = json.Marshal(removeAds(data))
	if err != nil {
		return err
	}

	return common.Response(res, body)
}

func removeAds(data any) any {
	switch v := data.(type) {
	case []any:
		var filtered = make([]any, 0, len(v))
		for _, item := range v {
			if result := removeAds(item); result != nil {
				filtered = append(filtered, result)
			}
		}
		return filtered
	case map[string]any:
		if isNodeAds(v) {
			return nil
		}

		disableGilding(v)

		for key, value := range v {
			v[key] = removeAds(value)
		}
	}

	return data
}

func isAds(data map[string]any) bool {
	if data["adPayload"] != nil || data["adEvents"] != nil {
		return true
	}

	if v, ok := data["__typename"].(string); ok && v == "AdPost" {
		return true
	}

	return false
}

func isNodeAds(data map[string]any) bool {
	if node, ok := data["node"].(map[string]any); ok && isAds(node) {
		return true
	}

	return isAds(data)
}

func disableGilding(data map[string]any) {
	if isGildable, ok := data["isGildable"].(bool); ok && isGildable {
		data["isGildable"] = false
	}
}
