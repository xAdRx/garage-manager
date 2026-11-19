package main

import (
	"log"
	"os"

	"restapi/server"
)

const (
	certFile = "certs/cert.pem"
	keyFile  = "certs/key.pem"
)

//go:generate swag init --outputTypes json,yaml

// @title			REST API
// @version		1.0
// @description	Basic Go REST API.
// @schemes		http https
// @BasePath		/
func main() {
	cfg := server.Config{
		HTTPAddr:  ":" + envOr("PORT", "5000"),
		StaticDir: os.Getenv("STATIC_DIR"),
	}

	// HTTPS runs only when the certificate files exist. On Render there are
	// none, because Render provides HTTPS in front of the app.
	if fileExists(certFile) && fileExists(keyFile) {
		cfg.HTTPSAddr = ":" + envOr("HTTPS_PORT", "5001")
		cfg.CertFile = certFile
		cfg.KeyFile = keyFile
	} else {
		log.Printf("No %s / %s found, HTTPS is off", certFile, keyFile)
	}

	log.Fatal(server.New(cfg).Start())
}

// envOr returns the environment variable key, or fallback when it is empty.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func fileExists(name string) bool {
	info, err := os.Stat(name)
	return err == nil && !info.IsDir()
}
