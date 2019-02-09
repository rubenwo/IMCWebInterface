package main

import (
	"IMCWebInterface/API"
	"IMCWebInterface/FrontEnd"
	"IMCWebInterface/IMC"
)

func main() {
	imc := IMC.TcpIMCServer{Port: "10000"}
	webServer := FrontEnd.Server{}
	apiServer := API.Server{}
	go imc.Start()
	go apiServer.Start(&imc)
	webServer.Start()
}
