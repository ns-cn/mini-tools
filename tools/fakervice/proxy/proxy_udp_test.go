package proxy

import (
	"fakervice/constants/signal"
	"fakervice/proxy/pool"
	"fmt"
	"testing"
	"time"
)

func TestUdpProxy_Proxy(t *testing.T) {
	udp := UdpProxy{}
	go func() {
		<-time.After(10 * time.Second)
		pool.Update(8847, signal.QUIT)
	}()
	err := udp.Proxy(8847, "127.0.0.1:8848")
	if err != nil {
		fmt.Println(err)
		return
	}
}
