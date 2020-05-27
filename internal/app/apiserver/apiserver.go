package apiserver

import (
	"github.com/alonelegion/golang_rest_api/internal/app/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *database.Database
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureDatabase(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting API server ...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/summary", s.HandleSummary).Methods("GET")
	s.router.HandleFunc("/positions", s.HandlePositions).Methods("GET")
	s.router.HandleFunc("/hello", s.helloPage).Methods("GET")
}

func (s *APIServer) configureDatabase() error {
	st := database.New(s.config.Database)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) helloPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Summary")
}
