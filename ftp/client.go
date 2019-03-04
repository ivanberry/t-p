package main

import (
	"bufio"
	"bytes"
	"fakeHttp/util"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	uiDir = "dir"
	uiCd = "cd"
	uiPwd = "pwd"
	uiQuit = "quit"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	host := os.Args[1]

	// establish tcp connection
	conn, err := net.Dial("tcp", host+":1202")
	util.CheckError(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, " \t\r\n")
		if err != nil {
			break
		}

		// split into command + arg
		strs := strings.SplitN(line, " ", 2)

		switch strs[0] {
		case uiDir:
			dirRequest(conn)
		case uiCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("cd \"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case uiPwd:
			pwdRequest(conn)
		case uiQuit:
			conn.Close()
		default:
			fmt.Println("Unknow command")
		}
	}
}

func dirRequest(conn net.Conn)  {
	conn.Write([]byte(util.DIR+" "))

	var buf [512]byte
	result := bytes.NewBuffer(nil)

	for {
		// read till hit a blank line

		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0: length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string)  {
	conn.Write([]byte(util.CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("Failed to change dir")
	}
}

func pwdRequest(conn net.Conn)  {
	conn.Write([]byte(util.PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir \"" +s + "\"")
}

