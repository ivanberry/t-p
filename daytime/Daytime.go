package main

import (
	"fakeHttp/util"
	"net"
	"time"
)

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	util.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	util.CheckError(err)

	//endless loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
