package common

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	dnsTimeout   = 30 * time.Second
	dnsKeepAlive = 30 * time.Second
	dnsNetwork   = "udp"
	dnsAddress   = "1.1.1.1:53"
)

func SetupCustomResolver() {
	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout:   dnsTimeout,
					KeepAlive: dnsKeepAlive,
				}

				return d.DialContext(ctx, dnsNetwork, dnsAddress)
			},
		},
	}

	http.DefaultTransport.(*http.Transport).DialContext = dialer.DialContext
}
