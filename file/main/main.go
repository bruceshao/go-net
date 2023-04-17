/*
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fmt.Println("read start")
	s := read2()
	fmt.Println(s)
	time.Sleep(10 * time.Minute)
}

func read1() string {
	fd, err := os.Open("/root/my.txt")
	if err != nil {
		panic(err)
	}
	//defer fd.Close()
	bz := make([]byte, 1024)
	n, err := fd.Read(bz)
	if err != nil {
		panic(err)
	}
	fd.Close()
	return string(bz[:n])
}

func read2() string {
	bz, err := os.ReadFile("/root/my.txt")
	if err != nil {
		panic(err)
	}
	return string(bz)
}

func read3() string {
	bz, err := ioutil.ReadFile("/root/my.txt")
	if err != nil {
		panic(err)
	}
	return string(bz)
}
