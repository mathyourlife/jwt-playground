package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var mySigningKey = []byte("secret")

type MyCustomClaims struct {
	jwt.StandardClaims
	UserID int64  `json:"user_id"`
	Roles  []Role `json:"roles"`
}

var PostLoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var username, password string
	var ok bool

	if username, password, ok = r.BasicAuth(); !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
		return
	}

	user, err := checkPassword(username, password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
		return
	}
	log.Printf("logging in user_id: %#v\n", user)

	// Create our custom claims instance
	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
		user.UserID,
		getUserRoles(user.UserID),
	}

	/* Create the token */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})

func checkPassword(username, password string) (*User, error) {
	for _, user := range Users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("failed login for username: %s", username)
}

func getUserRoles(userID int64) []Role {
	return []Role{Admin}
}

func loggedInHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &MyCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource")
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
			return
		}

		log.Printf("%#v\n", token.Claims.(*MyCustomClaims))
		h(w, r)

	}
}
