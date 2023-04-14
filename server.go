/*
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package net

import (
	"syscall"
)

const (
	BackLog = 128
)

type Server struct {
	IpAddr    [4]byte
	Port      int
	BackLog   int
	ReuseAddr bool
	ReusePort bool
	fd        int
}

func (s *Server) Listen() error {
	var (
		reuseAddr = 0
		reusePort = 0
		backLog   = BackLog
	)
	if s.ReuseAddr {
		reuseAddr = 1
	}
	if s.ReusePort {
		reusePort = 1
	}
	if s.BackLog > 0 {
		backLog = s.BackLog
	}
	sfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
	if err != nil {
		return err
	}
	err = syscall.SetsockoptInt(sfd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, reuseAddr)
	if err != nil {
		return err
	}
	err = syscall.SetsockoptInt(sfd, syscall.SOL_SOCKET, syscall.SO_REUSEPORT, reusePort)
	if err != nil {
		return err
	}
	err = syscall.Bind(sfd, &syscall.SockaddrInet4{
		Addr: s.IpAddr,
		Port: s.Port,
	})
	if err != nil {
		return err
	}
	err = syscall.Listen(sfd, backLog)
	if err != nil {
		return err
	}
	s.fd = sfd
	return nil
}

func (s *Server) Accept() (*Conn, error) {
	cfd, _, err := syscall.Accept(s.fd)
	if err != nil {
		return nil, err
	}
	return newConn(cfd), nil
}

func (s *Server) Close() error {
	return syscall.Close(s.fd)
}
