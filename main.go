package main

import (
	"github.com/labac-dev/surge-proxy/internal/busuu"
	"github.com/labac-dev/surge-proxy/internal/common"
	"github.com/labac-dev/surge-proxy/internal/proxy"
	"github.com/labac-dev/surge-proxy/internal/reddit"
	"github.com/labac-dev/surge-proxy/internal/truecaller"
)

func main() {
	app := proxy.NewProxy()

	app.OnDomain("gql-fed.reddit.com", reddit.ModifyResponse)
	app.OnDomain("premium-noneu.truecaller.com", truecaller.ModifyResponse)
	app.OnDomain("api.busuu.com", busuu.ModifyResponse)

	common.Serve(":https", app)
}
