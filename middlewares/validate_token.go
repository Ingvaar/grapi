package middlewares

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"

	"grapi/core"
	"grapi/utils"
)

// ValidateMiddleware : Validate the JWT
func ValidateMiddleware(config core.Config, level int, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if level > 0 {
			authorizationHeader := r.Header.Get("authorization")
			if authorizationHeader != "" {
				token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
					_, ok := token.Method.(*jwt.SigningMethodHMAC)
					if !ok {
						return nil, errors.New("Internal error")
					}
					return []byte(config.Secret), nil
				})
				if err != nil {
					utils.SendError(w, err, http.StatusInternalServerError)
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					uLevel := token.Claims.(jwt.MapClaims)["level"].(float64)
					if int(uLevel) >= level {
						next(w, r)
					} else {
						utils.SendError(w, errors.New("Unauthorized"), http.StatusUnauthorized)
						return
					}
				} else {
					utils.SendError(w, errors.New("Bad token"), http.StatusBadRequest)
					return
				}
			} else {
				utils.SendError(w, errors.New("Unauthorized"), http.StatusUnauthorized)
				return
			}
		} else {
			next(w, r)
		}
	})
}
