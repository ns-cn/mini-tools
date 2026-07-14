package consul

import (
	"github.com/hashicorp/consul/api"
	"sync"
)

// 连接池实例
var pool connectPool

// 链接池
type connectPool struct {
	mapper sync.Map
}

// Get 获取具体配置的链接实例
func (pool *connectPool) Get(config *api.Config) (*api.Client, error) {
	// TODO 链接复用待实现
	return api.NewClient(config)
}
