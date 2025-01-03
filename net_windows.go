package main

import (
	"net"
	"syscall"
	"golang.org/x/sys/windows"
)

func init() {
	listenConfig = net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
		    var opErr error
		    if err := c.Control(func(fd uintptr) {
		        opErr = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
		    }); err != nil {
		        return err
		    }
		    return opErr
		},
	}
}
