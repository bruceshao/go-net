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
	"time"
)

func main() {
	fmt.Println("read start")
	s := read()
	fmt.Println(s)
	time.Sleep(time.Hour)
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
