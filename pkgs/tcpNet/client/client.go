package client

import (
	"encoding/gob"
	"io"
	"net"
)

type Client struct {
	Conn *net.Conn
	Addr string
}

func NewClient(addr string) (Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return Client{}, err
	}
	return Client{Conn: &conn, Addr: addr}, nil

}

func GetClients(workerPorts []string) ([]Client, error) {
	var clients []Client
	for _, v := range workerPorts {
		client, err := NewClient("127.0.0.1:" + v)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func SendPayload(client Client, payload interface{}) error {
	encoder := gob.NewEncoder(*client.Conn)
	return encoder.Encode(payload)
}

func ReadResponse(client Client, payload interface{}) {
	var err error = nil
	for err == nil {
		dec := gob.NewDecoder(*client.Conn)
		err = dec.Decode(payload)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
}
