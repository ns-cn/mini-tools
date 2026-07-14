package proxy

import (
	"fakervice/constants/protocol"
	"fakervice/constants/signal"
	"fakervice/proxy/pool"
	"fmt"
	"io"
	"log"
	"net"
)

// TcpProxy TCP的代理
type TcpProxy struct {
	Proxyable // 可代理的
}

func (proxy TcpProxy) Proxy(port int, remote string) error {
	control := make(chan signal.Control, 1)
	err := pool.New(port, protocol.TCP, control)
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)
	pause := false
	for {
		select {
		case controlSignal := <-control:
			switch controlSignal {
			case signal.PAUSE:
			case signal.QUIT:
			case signal.CONTINUE:
				info, ok := pool.GetInfo(port)
				if !ok {
					break
				}
				if info.CanDo(controlSignal) {
					if controlSignal == signal.CONTINUE {
						pause = false
					} else if controlSignal == signal.PAUSE {
						pause = true
					} else if controlSignal == signal.QUIT {
						break
					}
					pool.Update(port, controlSignal)
				}
			default:
				if !pause {
					conn, err := listener.Accept()
					if err != nil {
						fmt.Printf("建立连接错误:%v\n", err)
						continue
					}
					go func(connection net.Conn) {
						defer func(connection net.Conn) {
							_ = connection.Close()
						}(connection)
						proxyConnection, err := net.Dial("tcp", remote)
						if err != nil {
							fmt.Printf("连接%v失败:%v\n", remote, err)
							return
						}
						ExitChan := make(chan bool, 2)
						proxyTo := func(source net.Conn, dest net.Conn, Exit chan bool) {
							_, err := io.Copy(dest, source)
							if err != nil && err != io.EOF {
								fmt.Printf("往%v发送数据失败:%v\n", remote, err)
							}
							ExitChan <- true
						}
						// 转发请求
						go proxyTo(connection, proxyConnection, ExitChan)
						go proxyTo(proxyConnection, connection, ExitChan)
						<-ExitChan
						<-ExitChan
						_ = proxyConnection.Close()
					}(conn)
				}
			}
		}
	}
	return nil
}
