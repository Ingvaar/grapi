package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"grapi/core"
)

// Logger : Prints logs of the server
func Logger(inner http.Handler, name string, config core.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		log.SetOutput(os.Stdout)
		log.Printf("%s\t%s\t%s\t%s\t%s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))

		if config.LogFile != "" {
			f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Printf("Error opening file: %v", err)
			}
			defer f.Close()
			log.SetOutput(f)
			log.Printf("%s\t%s\t%s\t%s\t%s",
				r.RemoteAddr,
				r.Method,
				r.RequestURI,
				name,
				time.Since(start))
		}
	})
}
