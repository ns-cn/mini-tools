package proxy

// Proxyable 可代理的
type Proxyable interface {
	// Proxy 监听指定端口,代理到目标地址
	Proxy(int, string) error
}
