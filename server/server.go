package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	c "grapi/core"
	m "grapi/middlewares"
)

var httpPort string
var httpsPort string
var address string
var certsDir string

// Server :
type Server struct {
	config *c.Config
	router *mux.Router
}

// Start :
func (s *Server) Start(router *mux.Router, config *c.Config) {
	s.router = router
	s.config = config
	s.checkConfig()
	httpPort = ":" + s.config.ServerPort
	address = s.config.ServerAddress
	certsDir = s.config.CertsDir
	cert := certsDir + "/cert.pem"
	key := certsDir + "/key.pem"
	s.checkPorts()

	if s.config.HTTPS != 0 {
		_, err := os.Stat(cert)
		_, err2 := os.Stat(key)
		if err != nil || err2 != nil {
			log.Fatal("Cert files not found")
		}
		if s.config.HTTPSOnly != 0 {
			log.Printf("Http server at %v%v redirecting to %v%v", address, httpPort, address, httpsPort)
			go s.loggedRedirectServer()
		} else {
			log.Printf("Http server started at %v%v", address, httpPort)
			go func() { log.Fatal(http.ListenAndServe(address+httpPort, s.router)) }()
		}
		log.Printf("Https server started at %v%v", address, httpsPort)
		log.Fatal(http.ListenAndServeTLS(httpsPort, cert, key, s.router))
	} else {
		log.Printf("Http server started at %v%v", address, httpPort)
		log.Fatal(http.ListenAndServe(address+httpPort, s.router))
	}
}

func (s *Server) loggedRedirectServer() {
	log.Fatal(http.ListenAndServe(address+httpPort, m.Logger(http.HandlerFunc(s.redirectToHTTPS), "Redirect", *s.config)))
}

func (s *Server) redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+address+httpsPort+r.RequestURI, http.StatusMovedPermanently)
}

func (s *Server) checkConfig() {
	if s.config.ServerAddress == "" {
		log.Fatal("Missing server address in config file")
	}
	if s.config.ServerPort == "" {
		s.config.ServerPort = "8080"
	}
	if s.config.CertsDir == "" {
		s.config.CertsDir = "."
	}
}

func (s *Server) checkPorts() {
	port, err := strconv.Atoi(s.config.ServerPort)

	if err != nil {
		httpPort = ":8080"
		port = 8080
	}
	httpsPort = strconv.Itoa(port + 1)
	httpsPort = ":" + httpsPort
}
