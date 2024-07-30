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
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		log.Fatalf("Failed to establish connection. Error: %s", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to establish connection. Error: %s", err)
		}
	}
}
