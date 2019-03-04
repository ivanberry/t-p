package main

import (
	"fakeHttp/util"
	"net"
)

func main() {

	// make a available tcp address
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	util.CheckError(err)

	// init a listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	util.CheckError(err)

	// endless loop to Accept
	for {
		conn, err := listener.Accept()
		// ignore the error
		if err != nil {
			continue
		}

		// run as goroutine
		go util.HandleClient(conn)
	}
}