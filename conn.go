/*
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package net

import (
	"fmt"
	"syscall"
)

type Conn struct {
	cfd int
}

func newConn(cfd int) *Conn {
	return &Conn{
		cfd: cfd,
	}
}

func (c *Conn) Close() error {
	return syscall.Close(c.cfd)
}

func (c *Conn) Read() ([]byte, error) {
	buf := make([]byte, 0)
	for {
		innerBuf := make([]byte, 1024)
		n, err := syscall.Read(c.cfd, innerBuf)
		if err != nil {
			return nil, err
		}
		buf = append(buf, innerBuf[0:n]...)
		fmt.Println("read ", n)
		if n < 1024 {
			// 表示读取完成
			return buf, nil
		}
	}
}
