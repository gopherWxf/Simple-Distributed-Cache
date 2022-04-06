package http

import (
	"bytes"
	"net/http"
)

type rebalanceHandler struct {
	*Server
}

func (h *rebalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	go h.rebalance()
}

func (h *rebalanceHandler) rebalance() {
	s := h.NewScanner()
	defer s.Close()
	client := &http.Client{}
	for s.Scan() {
		k := s.Key()
		redirectAddr, ok := h.ShouldProcess(k)
		if !ok {
			r, _ := http.NewRequest(http.MethodPut, "http://"+redirectAddr+":12345/cache/"+k, bytes.NewReader(s.Value()))
			client.Do(r)
			h.Del(k)
		}
	}
}

func (s *Server) rebalanceHandler() http.Handler {
	return &rebalanceHandler{s}
}
