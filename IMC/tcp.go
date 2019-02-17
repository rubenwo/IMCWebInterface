package IMC

import (
	"IMCWebInterface/API"
	"fmt"
	"net"
)

type Connection struct {
	Connected    bool
	Conn         net.Conn
	LatestConfig API.Config
}

var imcMap map[string]Connection

type TcpIMCServer struct {
	Port    string
	Running bool
}

func (s *TcpIMCServer) Start() error {
	s.Running = true
	imcMap = make(map[string]Connection)
	ln, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}
	fmt.Println("IMC Server is running...")
	for {
		conn, err := ln.Accept()

		if err != nil {
			return err
		}
		go onConnect(conn)
	}
}

func (s *TcpIMCServer) AddConfig(radioID string, config *API.Config) {
	conn, ok := imcMap[radioID]
	if !ok {
		return
	}
	conn.LatestConfig = *config
	if !conn.Connected {
		return
	}
	go uploadConfig(&conn)
}

func onConnect(conn net.Conn) {
	id := "random string"
	config := requestConfig(conn)
	imcMap[id] = Connection{
		Connected:    true,
		Conn:         conn,
		LatestConfig: *config}

	for {
		var buff []byte
		conn.Read(buff)
		decode(buff)
	}
}

func decode(buff []byte) {

}

func onDisconnect(connection *Connection) {

}

func requestConfig(conn net.Conn) *API.Config {
	return &API.Config{
		RadioSettings: API.RadioSettings{

		},
		AlarmSettings: API.AlarmSettings{

		},
		GeneralSettings: API.GeneralSettings{

		}}
}

func uploadConfig(connection *Connection) {

}
