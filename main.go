package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "server host")
	port := flag.String("port", "8080", "server port")
	filename := flag.String("data", "request.json", "request file")
	flag.Parse()
	tcpServer, err := net.ResolveTCPAddr("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: ", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		fmt.Println("Dial failed: ", err.Error())
		os.Exit(1)
	}
	cwd, _ := os.Getwd()
	file, err := os.ReadFile(cwd + "/" + *filename)
	if err != nil {
		fmt.Println("File reading filed: ", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(file)
	if err != nil {
		fmt.Println("Write data failed: ", err.Error())
		os.Exit(1)
	}

	received := make([]byte, 1<<10)
	_, err = conn.Read(received)
	if err != nil {
		fmt.Println("Read data failed: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(received))

	conn.Close()
}
