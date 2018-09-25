package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	c "grapi/config"
	r "grapi/router"
)

var httpPort string
var httpsPort string
var address string
var certsDir string

// StartServer :
func StartServer() {
	checkConfig()
	httpPort = ":" + c.Cfg.ServerPort
	address = c.Cfg.ServerAddress
	certsDir = c.Cfg.CertsDir
	cert := certsDir + "/cert.pem"
	key := certsDir + "/key.pem"
	checkPorts()

	if c.Cfg.HTTPS != 0 {
		_, err := os.Stat(cert)
		_, err2 := os.Stat(key)
		if err != nil || err2 != nil {
			log.Fatal("No cert files found")
		}
		if c.Cfg.HTTPSOnly != 0 {
			log.Printf("Http redirecting on %v%v", address, httpPort)
			go log.Fatal(http.ListenAndServe(httpPort, http.HandlerFunc(redirectToHTTPS)))
		} else {
			log.Fatal(http.ListenAndServe(httpPort, r.Router))
		}
		log.Printf("Https server started at %v%v", address, httpsPort)
		log.Fatal(http.ListenAndServeTLS(httpsPort, cert, key, r.Router))
	} else {
		log.Printf("Http server started at %v%v", address, httpPort)
		log.Fatal(http.ListenAndServe(httpPort, r.Router))
	}
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+address+httpPort+r.RequestURI, http.StatusMovedPermanently)
}

func checkConfig() {
	if c.Cfg.ServerAddress == "" {
		log.Fatal("Missing server address in config file")
	}
	if c.Cfg.ServerPort == "" {
		c.Cfg.ServerPort = ":8080"
	}
	if c.Cfg.CertsDir == "" {
		c.Cfg.CertsDir = "."
	}
}

func checkPorts() {
	port, err := strconv.Atoi(c.Cfg.ServerPort)

	if err != nil {
		httpPort = ":8080"
		port = 8080
	}
	httpsPort = strconv.Itoa(port + 1)
}
