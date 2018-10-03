package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	c "grapi/config"
	"grapi/db"
	"grapi/utils"
)

type auth struct {
	id    int
	user  string
	pass  string
	level int
}

// JwtToken : jwt struct
type JwtToken struct {
	Token string `json:"token"`
}

// Login :
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	authStruct, err := checkCredentials(r.Form)
	if err != nil {
		utils.SendResponse(w, err, http.StatusBadRequest)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       authStruct.id,
		"username": authStruct.user,
		"password": authStruct.pass,
		"level":    authStruct.level,
	})
	tokenString, err := token.SignedString([]byte(c.Cfg.Secret))
	if err != nil {
		utils.SendResponse(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func checkCredentials(form url.Values) (*auth, error) {
	var username string
	var pass []byte
	authStruct := new(auth)

	if form["username"] != nil && form["password"] != nil {
		username = form["username"][0]
		pass = []byte(form["password"][0])
	} else {
		err := errors.New("Bad request")
		return nil, err
	}
	statement := "SELECT id, " + c.Cfg.AuthUserField +
		", " + c.Cfg.AuthPassField +
		", " + c.Cfg.AuthLevel +
		" FROM " + c.Cfg.AuthTable +
		" WHERE " + c.Cfg.AuthUserField + "=\"" + username + "\""
	rows, err := db.SQL.Query(statement)
	if err != nil {
		err = errors.New("Database Error")
		return nil, err
	}
	defer rows.Close()
	result, err := utils.RowsToMap(rows)
	if err != nil {
		return nil, err
	}
	hash := result["password"].([]uint8)
	err = bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		err = errors.New("Wrong password")
		return nil, err
	}
	authStruct.user = username
	authStruct.pass = string(hash)
	authStruct.level = utils.RContToInt(result["level"].([]uint8))
	authStruct.id = utils.RContToInt(result["id"].([]uint8))
	return authStruct, err
}
