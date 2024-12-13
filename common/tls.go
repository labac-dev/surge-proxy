package common

import (
	"crypto/tls"
	"fmt"
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

func GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
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
