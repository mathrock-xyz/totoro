package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net"
	"os"
)

type ErrMsg struct {
	Error string `json:"error"`
}

func handle(conn net.Conn) {
	defer conn.Close()

	buff := []byte{}
	if _, err := conn.Read(buff); err != nil {
		senderr(conn, err)
		return
	}

	hash := sha256.Sum256(buff)
	name := hex.EncodeToString(hash[:])
	name = data + "/" + name

	file, err := os.Create(name)
	if err != nil {
		senderr(conn, err)
		return
	}

	defer file.Close()

	if _, err := file.Write(buff); err != nil {
		senderr(conn, err)
		return
	}
}

func senderr(conn net.Conn, err error) {
	var errmsg ErrMsg

	msg, err := json.Marshal(errmsg)
	if err != nil {
		log.Error(err.Error())
		return
	}

	if _, err := conn.Write(msg); err != nil {
		log.Error(err.Error())
		return
	}
}
