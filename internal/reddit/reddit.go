package reddit

import (
	"encoding/json"
	"net/http"

	"github.com/labac-dev/surge-proxy/internal/common"
	"github.com/itchyny/gojq"
)

const QueryString = `del(.. | select(
	.node?.adPayload?.adEvents or
	.node?.adEvents or
	.adEvents or
	.__typename == "AdPost"
)) | walk(
	if type == "object" and has("isGildable") then
		.isGildable = false
	else
		.
	end
)`

var code *gojq.Code

func init() {
	var (
		err   error
		query *gojq.Query
	)

	query, err = gojq.Parse(QueryString)
	if err != nil {
		panic(err)
	}

	code, err = gojq.Compile(query)
	if err != nil {
		panic(err)
	}
}

// ModifyResponse modifies the response from reddit
func ModifyResponse(res *http.Response) error {
	var (
		data interface{}
		err  error
		body []byte
	)

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	iter := code.Run(data)
	v, _ := iter.Next()

	if err, ok := v.(error); ok {
		return err
	}

	body, err = gojq.Marshal(v)
	if err != nil {
		return err
	}

	return common.Response(res, body)
}
