package pool

import (
	"fakervice/constants/protocol"
	"fakervice/constants/signal"
	"fakervice/constants/status"
	"fmt"
	"sync"
)

// 声明实例
var pool proxyPool

// proxyPool 代理池
type proxyPool struct {
	mapper sync.Map
}

// New 针对指定端口重新声明监听协议和控制的管道
func New(port int, protocol protocol.Protocol, control chan signal.Control) error {
	inUse, ok := GetInfo(port)
	if ok && inUse.Status == status.RUNNING {
		return fmt.Errorf("port %d in use with protocol %v", inUse.Port, inUse.Protocol)
	}
	info := &PortInfo{Port: port, Protocol: protocol, Status: status.INIT, Control: control}
	pool.mapper.Store(port, info)
	return nil
}

// GetInfo 获取指定端口的代理监听情况
func GetInfo(port int) (*PortInfo, bool) {
	value, ok := pool.mapper.Load(port)
	if ok {
		info, parsed := value.(*PortInfo)
		if parsed {
			return info, true
		}
	}
	return nil, false
}

// InUse 池中指定端口是否正在被占用
func InUse(port int) bool {
	info, ok := GetInfo(port)
	return ok && info.Status == status.RUNNING
}

// Update 更新代理池中指定端口的状态
func Update(port int, signal signal.Control) bool {
	info, ok := GetInfo(port)
	if ok && info.Status == status.RUNNING {
		info.Control <- signal
		pool.mapper.Delete(port)
	}
	return true
}
