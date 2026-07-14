package register

// Registry 注册中心
type Registry interface {
	// Register 注册某个代理服务
	Register() error
	// DeRegister 注销某个代理服务
	DeRegister() error
}
