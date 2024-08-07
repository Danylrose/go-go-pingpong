package multithreaded_server

import (
	"fmt"
	"log"
	"math/rand"
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
	defer listener.Close()
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
		go connectionHandler(conn) // goroutine for connections
	}
}

func connectionHandler(connection net.Conn) {
	defer connection.Close()
	var buf []byte
	var err error
	buf = make([]byte, 1024)

	for {
		_, err = connection.Read(buf) // multithreading var, handles incoming connections
		if err != nil {
			log.Fatalf("Connection failed. Error: %s", err)
		}
		// place holder code for ping-pong
		distributedCoinToss(connection)
	}

}

func distributedCoinToss(rdwr net.Conn) {
	// new code to test out writing through the socket
	// have not tested it yet before push. (8/7/24 | 13:21)
	var coinToss = rand.Intn(2)
	var buf []byte = make([]byte, 1024)
	log.Println(string(buf))
	rdwr.Write([]byte("Flip a coin. Heads or tails."))
	//fmt.Printf("Received: %s", string(buf[:n]))
	if coinToss == 0 {
		rdwr.Write([]byte("Coin landed on tails."))
	} else {
		rdwr.Write([]byte("Coin landed on heads."))
	}
}
