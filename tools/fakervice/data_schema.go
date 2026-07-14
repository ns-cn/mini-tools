package main

import (
	"fakervice/constants/protocol"
	status2 "fakervice/constants/status"
	"fakervice/proxy"
	"fakervice/register"
)

type Schema struct {
	Name      string              // 名称
	Ports     int                 // 端口（拆分后）
	To        string              // 目标转发地址
	Protocol  protocol.Protocol   // 协议类型
	status    status2.Status      // 当前方案的状态
	proxy     proxy.Proxyable     // 所有具体代理
	registers []register.Registry // 所有的注册中心
}
