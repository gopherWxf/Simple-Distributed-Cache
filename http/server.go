package http

import (
	"Simple-Distributed-Cache/cache"
	"Simple-Distributed-Cache/cluster"
	"net/http"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.Handle("/cluster", s.clusterHandler())
	http.Handle("/rebalance", s.rebalanceHandler())
	http.ListenAndServe(s.Addr()+":12345", nil)
}

func New(cache cache.Cache, node cluster.Node) *Server {
	return &Server{cache, node}
}
