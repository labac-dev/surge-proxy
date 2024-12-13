package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/labac-dev/surge-proxy/common"
	"github.com/labac-dev/surge-proxy/reddit"
)

func main() {
	common.SetupCustomResolver()

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "https"
			req.URL.Host = req.Host

			req.Header.Del("Accept-Encoding")
		},
		ModifyResponse: func(res *http.Response) error {
			switch res.Request.URL.Host {
			case "gql-fed.reddit.com":
				return reddit.ModifyResponse(res)
			default:
				return nil
			}
		},
	}

	common.Serve(":https", proxy)
}
