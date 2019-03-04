package main

import (
	"fakeHttp/util"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
	}

	// prepare ip address
	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Resolution error", err.Error())
	}
	util.CheckError(err)

	// prepare icmp conection
	conn, err := net.DialIP("ip4:icmp", addr, addr)
	util.CheckError(err)

	var msg [512]byte

	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37

	len := 8

	check := util.CheckSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	util.CheckError(err)

	_, err = conn.Read(msg[0:])
	util.CheckError(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}

	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}