package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/seriouspoop/GoAPI/types"
	"github.com/seriouspoop/GoAPI/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logged in"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetUserByEmail(r.Context(), payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest,
			fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	h.store.CreateUser(r.Context(), types.User{
		ID:        primitive.NewObjectID(),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
	})
}
