package server

import (
	"caring-task-1/pkgs/work"
	"encoding/gob"
	"fmt"
	"net"
	"sort"
)

type Server struct {
	Conn *net.Conn
}

func NewWorker(port string) error {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go func(c net.Conn) {
			handleRequest(c)
			err := c.Close()
			if err != nil {
				fmt.Print(err)
			}
		}(conn)
	}
}

func handleRequest(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	payload := work.WorkPayload{}
	dec.Decode(&payload)
	sort.Strings(payload.Data)
	sendResponse(&conn, payload)
}

func sendResponse(conn *net.Conn, payload work.WorkPayload) {
	encoder := gob.NewEncoder(*conn)
	encoder.Encode(payload)
}
