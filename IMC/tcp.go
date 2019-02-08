package IMC

import "net"

type TcpIMCServer struct {
	Port string
}

func (s *TcpIMCServer) Start() error {
	ln, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {

}
