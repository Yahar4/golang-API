package api

import (
	"database/sql"
	"gihub.com/Yahar4/service/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// APIServer this is a type of our server so we can perfectly connect to it
type APIServer struct {
	Addr string
	DB   *sql.DB
}

// NewAPIServer is a fucntion to create a server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		Addr: addr,
		DB:   db,
	}
}

// Run is for running our server with gorillaMux router
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("listening on address ", s.Addr)

	// server starts (you see this comment when do go run cmd/main.go)
	return http.ListenAndServe(s.Addr, router)
}
