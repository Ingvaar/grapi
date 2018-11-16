package sql

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

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
func (db *SQL) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	authStruct, err := db.checkCredentials(r.Form)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":	    authStruct.id,
		"username": authStruct.user,
		"password": authStruct.pass,
		"level":    authStruct.level,
	})
	tokenString, err := token.SignedString([]byte(db.config.Secret))
	if err != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func (db *SQL) checkCredentials(form url.Values) (*auth, error) {
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
	statement := "SELECT id, " + db.config.AuthUserField +
		", " + db.config.AuthPassField +
		", " + db.config.AuthLevel +
		" FROM " + db.config.AuthTable +
		" WHERE " + db.config.AuthUserField + "=\"" + username + "\""
	rows, err := db.DB.Query(statement)
	if err != nil {
		err = errors.New("Database Error")
		return nil, err
	}
	defer rows.Close()
	result, err := rowsToMap(rows)
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
