package consul

import (
	"fakervice/register"
	"github.com/hashicorp/consul/api"
)

// Instance consul的配置信息
type Instance struct {
	register.Registry
	UsingPool    bool                          // 是否使用池化链接共用链接,减少链接的创建
	Server       *api.Config                   // 目标服务器的注册信息
	Registration *api.AgentServiceRegistration // 注册信息
}

// GetConnection 通过配置创建与Consul的链接信息
func (instance *Instance) connect() (*api.Client, error) {
	if instance.UsingPool {
		return pool.Get(instance.Server)
	}
	// 如果不复用链接,仄每次重新创建链接
	return api.NewClient(instance.Server)
}

// Register 注册服务
func (instance *Instance) Register() error {
	connect, err := instance.connect()
	if err != nil {
		return err
	}
	return connect.Agent().ServiceRegister(instance.Registration)
}

// DeRegister 注销服务
func (instance *Instance) DeRegister() error {
	connect, err := instance.connect()
	if err != nil {
		return err
	}
	return connect.Agent().ServiceRegister(instance.Registration)
}
