package auth

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/mathyourlife/jwt-playground/pkg/authdb"
)

var testSecretKey = []byte("this is my secret key")

func TestBasicAuth(t *testing.T) {
	db := authdb.NewMockDB()

	db.User(&authdb.User{Username: "christie", Password: "c"}, authdb.OpCreate)
	db.User(&authdb.User{Username: "sally", Password: "s"}, authdb.OpCreate)
	db.User(&authdb.User{Username: "moxie", Password: "m"}, authdb.OpCreate)
	db.User(&authdb.User{Username: "susanne", Password: "s"}, authdb.OpCreate)
	db.User(&authdb.User{Username: "jack", Password: "j"}, authdb.OpCreate)

	a, _ := NewAuth(db, func() []byte {
		return testSecretKey
	})
	req := LoginRequest{
		Username: "moxie",
		Password: "m",
	}
	resp, err := a.Login(req)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp)
	s, _ := db.User(&authdb.User{UserID: 2}, authdb.OpRead)

	db.User(nil, 8)
	db.User(s, authdb.OpDelete)
	db.User(nil, 8)

	r, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", resp.Token))
	a.ReadTokenFromRequest(r)

}
