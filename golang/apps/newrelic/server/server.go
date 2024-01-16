package server

import (
	"net/http"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

// Server handler
func (s *Server) Handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.createHttpResponse(w, http.StatusOK, []byte("OK"))
}

func (s *Server) createHttpResponse(
	w http.ResponseWriter,
	statusCode int,
	body []byte,
) {
	w.WriteHeader(statusCode)
	w.Write(body)
}
