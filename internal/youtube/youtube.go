package youtube

import (
	"io"
	"log"
	"net/http"

	"github.com/labac-dev/surge-proxy/internal/common"
)

func ModifyResponse(res *http.Response) error {
	log.Printf("Modifying response for %s", res.Request.URL.String())

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return common.Response(res, body)
}
