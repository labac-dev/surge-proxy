package common

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
)

var (
	certPEM = "tmp/cert.pem"
	keyPEM  = "tmp/key.pem"
)

func Serve(addr string, handler http.Handler) {
	server := &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: &tls.Config{GetCertificate: getCertificate()},
	}

	log.Printf("Starting server on %s", server.Addr)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func getCertificate() func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		for _, dnsName := range x509Cert.DNSNames {
			if dnsName == hello.ServerName {
				return &cert, nil
			}
		}

		return nil, fmt.Errorf("unknown server name: %s", hello.ServerName)
	}
}
