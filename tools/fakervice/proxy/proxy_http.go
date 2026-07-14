package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// HttpProxy HTTP代理
type HttpProxy struct {
	Proxyable // 可代理的
}

func (proxy HttpProxy) Proxy(port int, remote string) error {
	mux := http.NewServeMux()
	redirect := func(w http.ResponseWriter, r *http.Request) {
		cli := &http.Client{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Print("io.ReadFull(r.Body, body) ", err.Error())
		}
		var schema = "http"
		if r.URL.Scheme != "" {
			schema = r.URL.Scheme
		}
		url := &url.URL{Host: remote, Scheme: schema, Path: r.URL.Path, RawQuery: r.URL.RawQuery}
		reqURL := url.String()
		fmt.Println(url)
		req, err := http.NewRequest(r.Method, reqURL, strings.NewReader(string(body)))
		if err != nil {
			fmt.Print("http.NewRequest ", err.Error())
			return
		}
		for k, v := range r.Header {
			if k == "Host" {
				req.Header.Set("Host", remote)
			} else {
				req.Header.Set(k, v[0])
			}
		}
		res, err := cli.Do(req)
		if err != nil {
			fmt.Print("cli.Do(req) ", err.Error())
			return
		}
		defer res.Body.Close()
		for k, v := range res.Header {
			w.Header().Set(k, v[0])
		}
		io.Copy(w, res.Body)
	}
	mux.HandleFunc("/", redirect)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
