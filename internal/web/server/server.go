package web

import (
	"github.com/maxlvl/beedrill/config"
	"github.com/maxlvl/beedrill/internal/loadtester"
	"github.com/maxlvl/beedrill/scenarios"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	lt         *loadtester.LoadTester
	config_    config.Config
}

func NewServer(addr string, lt *loadtester.LoadTester, config_ config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: addr,
		},
		lt:      lt,
		config_: config_,
	}
}

func (s *Server) Start() error {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../internal/web/js"))))
	http.HandleFunc("/", s.HandleIndex)
	http.HandleFunc("/api/v1/start", s.HandleStart)
	s.httpServer.ListenAndServe()
	return nil
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../internal/web/index.html")
}

func (s *Server) HandleStart(w http.ResponseWriter, r *http.Request) {
	var scenarios_array []scenarios.Scenario
	for _, s := range s.config_.Scenarios {
		switch s.Type {
		case "simple":
			scenarios_array = append(scenarios_array, &scenarios.SimpleScenario{URL: s.URL})
		case "complex":
			scenarios_array = append(
				scenarios_array,
				&scenarios.ComplexScenario{
					GetURL:  s.GetURL,
					PostURL: s.PostURL,
					Payload: s.Payload,
				})
		}
	}
	s.lt.Run(scenarios_array)

}
