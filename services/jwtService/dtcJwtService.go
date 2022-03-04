package dtcJwtServices

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	// "strconv"
  "time"

  b64 "encoding/base64"
	// "github.com/gorilla/mux"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
  Now int64 `json:"now"`
	jwt.StandardClaims
}

func createToken() string {
  // create token
  mySigningKey := []byte("AllYourBase")

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "foo": "bar",
		"module": "IMC,RCA,DLL",
    "now": time.Now().Unix(),
    "exp": time.Now().Add(time.Hour * (1)).Unix(),
  })
	tokenString, _ := token.SignedString(mySigningKey)
  fmt.Println(tokenString)

  akuaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ1NzcxNzYsImZvbyI6ImJhciIsIm5vdyI6MTYxNDgzNjM3Nn0.V5v7J7A4qichB4QRRqlloowuEkOV9bqWwY_tdOkRV04"

  uEnc := b64.URLEncoding.EncodeToString([]byte(akuaToken))
  fmt.Println(uEnc)
  // uDec, _ := b64.URLEncoding.DecodeString("eyJleHAiOjE2MTUwOTU0MTAsImZvbyI6ImJhciIsIm5vdyI6MTYxNDgzNjIxMH0")
  // fmt.Println(string(uDec))

  // parser
  token2, err2 := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token2.Claims.(*MyCustomClaims); ok && token2.Valid {
		fmt.Printf("%v %v %v", claims.Foo, claims.Now, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err2)
	}

  return tokenString
}

func Login(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

	syanticToken := r.Header.Get("syantic-token")
	fmt.Println("from header : " + syanticToken)

	syanticKeys := r.URL.Query()["syantic-key"]
	if syanticKeys != nil && len(syanticKeys[0]) > 0 {
		syanticKey := syanticKeys[0]
		fmt.Println("from param : " + syanticKey)
	}

  token := createToken()

  json.NewEncoder(w).Encode(token)
}
