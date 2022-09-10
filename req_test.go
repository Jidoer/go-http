package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/kataras/iris/v12"
)

func Test_Req(t *testing.T) {
	go start_teat_server(t)
	params := &url.Values{
		"Method": {"GET", "POST"},
		"IP":     {"127.0.0.1"},
	}
	headers := &http.Header{"Content-Type": {"application/json"}, "Accept": {"application/json"}, "User-Agent": {"Mozilla/5566, AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.299.252 Safari/537.36"}}
	s := NewSession()
	cookiesUrl, _ := url.Parse("http://127.0.0.1:5566/test")
	s.CookiesJar.SetCookies(cookiesUrl, []*http.Cookie{
		{Name: "foo", Value: "bar"},
		{Name: "bar", Value: "baz"},
		{Name: "baz", Value: "foo"},
	})
	s.Headers = headers
	r, e := s.Get("http://127.0.0.1:5566/test/", params)
	log.Println("Req Error:", e)

	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	log.Println("Body:", body)
}

func start_teat_server(t *testing.T) {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//INDEX
	app.Get("/test", func(ctx iris.Context) {
		log.Println(ctx.ResponseWriter())
		log.Println(ctx.Request())
		ctx.ResponseWriter().Write([]byte("Hello, world!"))
	})
	app.Run(iris.Addr("127.0.0.1:5566"))

}
