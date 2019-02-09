package IMC

import "net"

type TcpIMCServer struct {
	Port    string
	Running bool
}

func (s *TcpIMCServer) Start() error {
	s.Running = true
	ln, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

}
