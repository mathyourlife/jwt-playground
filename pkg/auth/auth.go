package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/mathyourlife/jwt-playground/pkg/authdb"
)

type Auth interface {
	Login(LoginRequest) (LoginResponse, error)
}

type GenSigningKey func() []byte

type auth struct {
	db     authdb.DB
	getKey GenSigningKey
}

func NewAuth(db authdb.DB, getKey GenSigningKey) (*auth, error) {
	return &auth{
		db:     db,
		getKey: getKey,
	}, nil
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	UserID int64
	Token  string
}

func (a *auth) Login(req LoginRequest) (*LoginResponse, error) {
	userID, err := a.db.BasicAuth(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	claims := UserCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
		userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.getKey())
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		UserID: userID,
		Token:  tokenString,
	}, nil
}

type UserCustomClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func (a *auth) ReadTokenFromRequest(r *http.Request) {

	t, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &UserCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return a.getKey(), nil
		})

	if err != nil {
		log.Println(err)
		return
	}

	if !t.Valid {
		log.Println("Token is not valid")
		return
	}

	log.Printf("%#v\n", t.Claims.(*UserCustomClaims))

}

// type Admin interface {
// 	CreateUser(string, *User) (*User, error)
// }

// type admin struct {
// 	db authdb.DB
// }

// func NewAdmin(db authdb.DB) (*admin, error) {
// 	return &admin{
// 		db: db,
// 	}
// }

// func (a *admin) CreateUser(u *User) (*User, error) {

// }
