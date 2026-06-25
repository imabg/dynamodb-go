package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	localConfig "github.com/imabg/dynamodb-go/config"
	"go.uber.org/zap"
)

type App struct {
	logger *zap.SugaredLogger
	router *mux.Router
	env localConfig.Config
}

type Server struct {
	app 	*App
	server *http.Server
}

func NewServer(env localConfig.Config, logger *zap.SugaredLogger) *Server {
	router := mux.NewRouter()
	app := &App{
		logger: logger,
		router: router,
		env: env,
	}
	return &Server{
		app: app,
		server: &http.Server{
			Addr:    env.SERVER_PORT,
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler: router,
		},
	}
}

func (s *Server) Start() {
	s.server.Addr = s.app.env.SERVER_PORT
	s.server.ReadHeaderTimeout = 5 * time.Second // 5 seconds
	s.app.logger.Infof("Server started on port %s", s.app.env.SERVER_PORT)
	s.app.logger.Fatal(s.server.ListenAndServe())
}

func (s *Server) Stop() {
	s.app.logger.Infof("Server stopped")
	s.app.logger.Fatal(s.server.Close())
}