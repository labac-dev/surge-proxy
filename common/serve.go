package common

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

var certSearchPaths = [][]string{
	{"/etc/ssl/private/cert.pem", "/etc/ssl/private/key.pem"},
	{"tmp/cert.pem", "tmp/key.pem"},
}

var (
	certPEM = []byte(os.Getenv("CERTIFICATE"))
	keyPEM  = []byte(os.Getenv("PRIVATE_KEY"))
)

func Serve(addr string, handler http.Handler) {
	server := &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: &tls.Config{GetCertificate: getCertificate},
	}

	log.Printf("Starting server on %s", server.Addr)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func getCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	for _, paths := range certSearchPaths {
		cert, err := tls.LoadX509KeyPair(paths[0], paths[1])
		if err == nil {
			return &cert, nil
		}
	}

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err == nil {
		return &cert, nil
	}

	return nil, fmt.Errorf("No certificate found")
}
