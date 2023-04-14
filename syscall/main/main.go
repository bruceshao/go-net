/*
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package main

import (
	"fmt"
	net "go-net"
)

func main() {
	c := make(chan struct{})
	server1 := &net.Server{
		IpAddr:    [4]byte{0, 0, 0, 0},
		Port:      9090,
		BackLog:   3,
		ReuseAddr: true,
		ReusePort: true,
	}
	go start(server1)
	server2 := &net.Server{
		IpAddr:    [4]byte{127, 0, 0, 1},
		Port:      9090,
		BackLog:   3,
		ReuseAddr: true,
		ReusePort: true,
	}
	go start(server2)
	<-c
}

func start(server *net.Server) {
	err := server.Listen()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			_ = server.Close()
		}
		go func(conn *net.Conn) {
			for {
				msg, err := conn.Read()
				if err != nil {
					_ = conn.Close()
				}
				fmt.Println(string(msg))
			}
		}(conn)
	}
}
