package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	// 监听80端口，响应所有请求并返回hello world
	http.ListenAndServe(":80",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello world")
			// 获取接口列表
			ifaces, err := net.Interfaces()
			if err != nil {
				panic(err)
			}
			// 迭代每个接口
			for _, iface := range ifaces {
				// 忽略非物理接口
				if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
					continue
				}
				// 检索该接口的硬件地址（MAC 地址）
				addr := iface.HardwareAddr
				if len(addr) > 0 {
					fmt.Fprintln(w, "MAC address:", addr.String())
				}
			}
		}))
}
