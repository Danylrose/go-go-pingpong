package multithreaded_server

import "net"

type Server struct {
	host string
	port int
}

type Config struct {
	Host string
	Port int
}

func distributedCoinToss(conn net.Conn) {
	// pushing
}
