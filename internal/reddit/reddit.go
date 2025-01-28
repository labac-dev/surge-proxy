package reddit

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/itchyny/gojq"
	"github.com/labac-dev/surge-proxy/internal/common"
)

const QueryString = `del(.. | select(
	.node?.adPayload?.adEvents or
	.node?.adEvents or
	.adEvents or
	.__typename == "AdPost"
))`

var query *gojq.Query

func init() {
	var err error

	query, err = gojq.Parse(QueryString)
	if err != nil {
		panic(err)
	}
}

func ModifyResponse(res *http.Response) error {
	var (
		data interface{}
		err  error
		body []byte
	)

	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	iter := query.Run(data)
	v, _ := iter.Next()

	if err, ok := v.(error); ok {
		return err
	}

	body, err = gojq.Marshal(v)
	if err != nil {
		return err
	}

	replacements := map[string]string{
		`"isGildable":true`:       `"isGildable":false`,
		`"isPremiumMember":false`: `"isPremiumMember":true`,
		`"isSubscribed":false`:    `"isSubscribed":true`,
		`"isEmployee":false`:      `"isEmployee":true`,
	}

	for old, new := range replacements {
		body = bytes.ReplaceAll(body, []byte(old), []byte(new))
	}

	return common.Response(res, body)
}
