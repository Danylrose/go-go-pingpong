package multithreaded_server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	host string
	port int
}

type Config struct {
	Host string
	Port int
}

func NewServer(config *Config) *Server {
	return &Server{host: config.Host, port: config.Port}
}

func (server *Server) RunServer() {
	//listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.host, server.port))
	var listener net.Listener
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", server.host, server.port))

	if err != nil {
		log.Fatalf("Cannot establish connection. Error: %v", err)
	}

	var conn net.Conn

	for {
		conn, err = listener.Accept()
		if err != nil {
			log.Fatalf("Connection failed. Error: %s", err)
		}

	}
}

func connectionHandler(connection net.Conn) {
	defer connection.Close()
	var buf []byte
	var numBytes int
	var err error
	buf = make([]byte, 1024)

	for {
		numBytes, err = connection.Read(buf)
		if err != nil {
			log.Fatalf("Connection failed. Error: %s", err)
		}
		// place holder code for ping-pong
	}

}

func distributedCoinToss() {

}
