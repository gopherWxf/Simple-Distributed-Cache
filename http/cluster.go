package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type clusterHandler struct {
	*Server
}

func (h *clusterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	m := h.Members()
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (s *Server) clusterHandler() http.Handler {
	return &clusterHandler{s}
}
