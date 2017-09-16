package server

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathan-osman/go-ldserver/manager"
	"github.com/nathan-osman/go-ldserver/presenter"
	"github.com/sirupsen/logrus"
)

// Server allows the user to start and stop performances.
type Server struct {
	listener  net.Listener
	server    *http.Server
	manager   manager.Manager
	presenter *presenter.Presenter
	log       *logrus.Entry
	stoppedCh chan bool
}

func (s *Server) run() {
	defer close(s.stoppedCh)
	defer s.log.Info("web server stopped")
	s.log.Info("web server started")
	if err := s.server.Serve(s.listener); err != nil {
		s.log.Error(err.Error())
	}
}

// NewServer creates a new web server.
func NewServer(cfg *Config, m manager.Manager) (*Server, error) {
	l, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return nil, err
	}
	var (
		router = mux.NewRouter()
		s      = &Server{
			listener: l,
			server: &http.Server{
				Handler: router,
			},
			manager:   m,
			log:       logrus.WithField("context", "server"),
			stoppedCh: make(chan bool),
		}
	)
	router.HandleFunc("/", s.index)
	go s.run()
	return s, nil
}

// Close shuts down the web server.
func (s *Server) Close() {
	s.listener.Close()
	<-s.stoppedCh
}
