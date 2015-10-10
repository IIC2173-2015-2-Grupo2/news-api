package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

var (
	user   = &User{Name: "DummyUser", Username: "username", Password: "password", Email: "email@email.com"}
	server *httptest.Server
)

func TestMain(m *testing.M) {
	server = httptest.NewServer(Server(nil))
	fmt.Println("Server URL", server.URL)
	defer server.Close()
	os.Exit(m.Run())
}

func TestIndex(t *testing.T) {
	Convey("Should be able to access index and redirect to APIv1 index", t, func() {
		res, _ := http.Get(server.URL)
		defer res.Body.Close()

		So(res.StatusCode, ShouldEqual, 200)
	})
}

func TestLogin(t *testing.T) {
	Convey("Should be able to login", t, func() {
		// jsondata, _ := json.Marshal(user)
		// data := strings.NewReader(string(jsondata))
		// req, _ := http.NewRequest("POST", server.URL+"/api/v1/auth/signup", data)
		// client := &http.Client{}
		// res, err := client.Do(req)
		// fmt.Println(res, err)
		// So(res.StatusCode, ShouldEqual, 200)
		// res, _ := http.PostForm(server.URL+"/api/v1/auth/signup", url.Values{
		// 	"username": {user.Username},
		// 	"password": {user.Password},
		// 	"name":     {user.Name},
		// 	"email":    {user.Email},
		// })
		// So(res.StatusCode, ShouldEqual, 200)
	})
}
