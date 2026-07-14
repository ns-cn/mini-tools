package pool

import (
	"fakervice/constants/protocol"
	"fakervice/constants/signal"
	"fakervice/constants/status"
)

// PortInfo 端口信息
type PortInfo struct {
	Port     int                 // 端口
	Protocol protocol.Protocol   // 监听类型
	Status   status.Status       // 状态
	Control  chan signal.Control // 控制信号量
}

// CanDo 当前状态能否进行如下的控制
func (info PortInfo) CanDo(control signal.Control) bool {
	switch info.Status {
	case status.INIT:
		return false
	case status.RUNNING:
		return control == signal.PAUSE || control == signal.QUIT
	case status.PAUSED:
		return control == signal.CONTINUE
	case status.TERMINATED:
		return control == signal.RESUME
	}
	return false
}
