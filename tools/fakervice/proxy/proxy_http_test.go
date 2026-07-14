package proxy

import (
	"fakervice/constants/signal"
	"fakervice/proxy/pool"
	"fmt"
	"testing"
	"time"
)

func TestHttpProxy_Proxy(t *testing.T) {
	http := HttpProxy{}
	go func() {
		<-time.After(10 * time.Second)
		pool.Update(8847, signal.QUIT)
	}()
	err := http.Proxy(8847, "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
}
