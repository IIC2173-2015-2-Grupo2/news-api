package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

const (
	url = "localhost:8000"
)

func TestIndex(t *testing.T) {
	server := httptest.NewServer(Server(nil))
	defer server.Close()

	Convey("Should be able to access index redirect to APIv1 index", t, func() {
		res, _ := http.Get(server.URL)
		So(res.StatusCode, ShouldEqual, 200)
	})
}
