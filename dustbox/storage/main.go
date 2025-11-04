package main

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/hashicorp/memberlist"
)

func main() {
	if err := init_storage(); err != nil {
		slog.Error(err.Error())
		return
	}

	if err := init_log(); err != nil {
		slog.Error(err.Error())
		return
	}

	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	_, err = list.Join([]string{""})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer listener.Close()

	log.Info("TCP Server listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		go handle(conn)
	}
}
