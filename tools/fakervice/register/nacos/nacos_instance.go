package nacos

import (
	"fakervice/register"
)

// Instance Nacos注册中心
type Instance struct {
	register.Registry //
}

// Register 注册某个代理服务
func (register Instance) Register() error {
	// TODO 注册到nacos注册中心待实现
	return nil
}

// DeRegister 注销某个代理服务
func (register Instance) DeRegister() error {
	// TODO 从nacos注册中心注销待实现
	return nil
}
