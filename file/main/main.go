/*
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	s := read()
	fmt.Println(s)
	c := make(chan struct{})
	<-c
}

func read() string {
	fd, err := os.Open("/proc/sys/net/core/somaxconn")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	bz := make([]byte, 1024)
	n, err := fd.Read(bz)
	if err != nil {
		panic(err)
	}
	return string(bz[:n])
}
