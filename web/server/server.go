package web

import (
	"github.com/maxlvl/gocust/internal/loadtester"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	lt         *loadtester.LoadTester
}

func NewServer(addr string, lt *loadtester.LoadTester) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: addr,
		},
		lt: lt,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.HandleIndex)
	// http.HandleFunc('/api/v1/start', s.HandleStart(),
	// http.HandleFunc('/api/v1/stop', s.HandleStop(),
	return s.httpServer.ListenAndServe()
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}
