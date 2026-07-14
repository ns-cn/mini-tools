package proxy

// HttpsProxy HTTPS代理
type HttpsProxy struct {
	Proxyable // 可代理的
}

func (proxy HttpsProxy) Proxy(port int) error {
	// TODO HTTPS代理待实现
	return nil
}
