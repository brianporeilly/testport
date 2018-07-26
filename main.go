package main

import (
	"os"
	"net"
	"fmt"
	"time"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("testport <server> <port> [<network>] [<timeout>]")
		fmt.Println("Defaults:")
		fmt.Println("network: tcp")
		fmt.Println("timeout: 5")
		os.Exit(255)
	}


	server := os.Args[1]
	port := os.Args[2]
	var seconds = 5
	var network = "tcp"
	if len(os.Args) > 3 {
		network = os.Args[3]
	}
	if len(os.Args) > 4 {
		seconds, _ = strconv.Atoi(os.Args[4])
	}

	_ , err := net.DialTimeout(network, server + ":" + port,time.Duration(seconds)*time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to " + server + ":" + port)
	os.Exit(0)

}
