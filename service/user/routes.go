package user

import (
	"gihub.com/Yahar4/types"
	"gihub.com/Yahar4/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler is a type for handling every function and endpoint
type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes for handling, you can see it in api.go
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// here are some endpoints functions
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if user exists

	// if it does we don't create new account
}
