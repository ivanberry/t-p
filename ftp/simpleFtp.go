package main

import (
	"fakeHttp/util"
	"fmt"
	"net"
	"os"
)


func main() {
	service := ":1202"

	// prepare address
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	util.CheckError(err)

	// prepare listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	util.CheckError(err)

	// endless loop for accept client connection
	for {
		fmt.Println("connection handle")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn)  {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(buf[0:n])

		// decode request
		if s[0:2] == util.CD {
			chdir(conn, s[3:])
		} else if s[0:3] == util.DIR {
			dirList(conn)
		} else if s[0:3] == util.PWD {
			pwd(conn)
		}
	}

}

// show dir content
func chdir(conn net.Conn, s string)  {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn)  {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func dirList(conn net.Conn)  {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}