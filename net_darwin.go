package main

import (
	"net"
	"syscall"
	"golang.org/x/sys/unix"
)

func init() {
	listenConfig = net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
		    var opErr error
		    if err := c.Control(func(fd uintptr) {
		        opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		    }); err != nil {
		        return err
		    }
		    return opErr
		},
	}
}
