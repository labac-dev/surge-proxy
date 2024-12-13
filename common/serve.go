package common

import (
	"crypto/tls"
	"log"
	"net/http"
)

func Serve(addr string, handler http.Handler) {
	server := &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: &tls.Config{GetCertificate: GetCertificate},
	}

	log.Printf("Starting server on %s", server.Addr)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}
