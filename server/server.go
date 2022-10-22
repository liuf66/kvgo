package server

import (
	"fmt"
	"log"
	"net"
	"syscall"

	"github.com/liuf66/kvgo/config"
)

func Start(c *config.ServerConfig) (err error) {
	serverSocketFd, err := createServerSocket(c.Host, c.Port)
	if err != nil {
		return
	}
	log.Printf("server listening at %v:%v\n", c.Host, c.Port)
	_ = serverSocketFd
	return
}

func createServerSocket(host string, port int) (serverSocketFd int, err error) {
	serverSocketFd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		err = fmt.Errorf("failed to create server socket: %v", err)
		return
	}

	err = syscall.SetsockoptInt(serverSocketFd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		err = fmt.Errorf("failed to set server socket reuseaddr: %v", err)
		return
	}

	err = syscall.SetNonblock(serverSocketFd, true)
	if err != nil {
		err = fmt.Errorf("failed to set server socket nonblock: %v", err)
		return
	}

	ipAddr := net.ParseIP(host)
	err = syscall.Bind(serverSocketFd, &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3]},
	})
	if err != nil {
		err = fmt.Errorf("failed to bind server socket: %v", err)
		return
	}

	err = syscall.Listen(serverSocketFd, 1024)
	if err != nil {
		err = fmt.Errorf("failed to listen server socket: %v", err)
		return
	}

	return
}
