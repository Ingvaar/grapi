package server

import (
	"crypto/tls"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	c "grapi/config"
	r "grapi/router"
)

// StartServer :
func StartServer() {
	port := ":" + c.Cfg.ServerPort
	address := c.Cfg.ServerAddress
	certsDir := c.Cfg.CertsDir

	if c.Cfg.HTTPS != 0 {
		checkConfig()
		certManager := autocert.Manager {
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(address),
			Cache:	    autocert.DirCache(certsDir),
		}

		server := &http.Server{
			Addr: port,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
		if c.Cfg.HTTPSOnly != 0 {
			go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
		}
		log.Printf("Server started at address %v on port %v", c.Cfg.ServerAddress, c.Cfg.ServerPort)
		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		log.Printf("Server started on port %v", c.Cfg.ServerPort)
		log.Fatal(http.ListenAndServe(port, r.Router))
	}
}

func checkConfig() {
	if c.Cfg.ServerAddress == "" {
		log.Fatal("Missing server address in config file")
	}
	if c.Cfg.ServerPort == "" {
		c.Cfg.ServerPort = ":https"
	}
	if c.Cfg.CertsDir == "" {
		c.Cfg.CertsDir = "."
	}
}
