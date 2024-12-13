package main

import (
	"github.com/labac-dev/surge-proxy/internal/common"
	"github.com/labac-dev/surge-proxy/internal/proxy"
	"github.com/labac-dev/surge-proxy/internal/reddit"
	"github.com/labac-dev/surge-proxy/internal/youtube"
)

func main() {
	common.SetupCustomResolver()

	app := proxy.NewProxy()
	app.OnDomain("gql-fed.reddit.com", reddit.ModifyResponse)
	app.OnDomain("youtubei.googleapis.com", youtube.ModifyResponse)

	common.Serve(":https", app)
}
