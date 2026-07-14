package main

import (
	"github.com/ns-cn/froute"
	"net/http"
)

func main() {
	frouters := new(froute.Frouters)
	group := froute.FrouteGroup{}
	check := func(floor int) func(string) bool {
		if floor == 1 {
			return func(s string) bool {
				return true
			}
		}
		return nil
	}
	group.SetMiddleware(func(writer http.ResponseWriter, request *http.Request, node *froute.FrouteNode) froute.MiddlewareResult {
		_, _ = writer.Write([]byte("processed in middleware!\n"))
		if node == nil {
			return froute.MiddlewareResult{Suspend: true}
		}
		return froute.MiddlewareResult{Suspend: false}
	})
	handler := func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello from handler!"))
	}
	group.AddHandler("GET", "/", check, handler)
	group.AddHandler("GET", "/user/info", check, handler)
	group.AddHandler("GET", "/user/card", check, handler)
	group.AddHandler("DELETE", "/user/mod", check, handler)
	frouters.AddGroup(&group)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		frouters.ServeHttp(writer, request)
	})
	_ = http.ListenAndServe(":8888", nil)
}