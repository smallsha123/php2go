package auth

import (
	json2 "encoding/json"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/smallsha123/php2go/globalkey"
	"github.com/tidwall/gjson"
)

func GenToken(secretKey string, iat, seconds int64, data map[string]interface{}) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["create_time"] = time.Now().Format(globalkey.DateTime)
	if data != nil {
		claims["jwtUserInfo"] = data
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	str, _ := token.SignedString([]byte(secretKey))
	return str
}

func GetUserId(authorization string, authSecret string) int64 {
	authorization = strings.Replace(authorization, "Bearer ", "", 1)
	if authorization == "" {
		return int64(0)
	}

	claims := make(jwt.MapClaims)
	userInfo, _ := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(authSecret), nil
	})

	if userInfo == nil {
		return int64(0)
	}
	json, _ := json2.Marshal(claims)
	userId := gjson.Get(string(json), "jwtUserId").Int()
	return userId
}
