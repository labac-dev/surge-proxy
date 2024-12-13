package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	reverseProxy *httputil.ReverseProxy
	middlewares  map[string]func(res *http.Response) error
}

func NewProxy() *Proxy {
	middlewares := make(map[string]func(res *http.Response) error)
	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "https"
			req.URL.Host = req.Host

			req.Header.Del("Accept-Encoding")
		},
		ModifyResponse: func(res *http.Response) error {
			if cb, ok := middlewares[res.Request.URL.Host]; ok {
				return cb(res)
			}

			return fmt.Errorf("No middleware found for %s", res.Request.URL.Host)
		},
	}

	return &Proxy{
		reverseProxy: reverseProxy,
		middlewares:  middlewares,
	}
}

func (p *Proxy) OnDomain(domain string, cb func(res *http.Response) error) {
	p.middlewares[domain] = cb
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.reverseProxy.ServeHTTP(w, r)
}
