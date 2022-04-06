package tcp

import (
	"Simple-Distributed-Cache/cache"
	"Simple-Distributed-Cache/cluster"
	"log"
	"net"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", s.Addr()+":12346")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go s.process(conn)
	}
}
func New(cache cache.Cache, node cluster.Node) *Server {
	return &Server{cache, node}
}
