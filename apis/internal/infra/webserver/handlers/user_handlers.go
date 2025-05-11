package handlers

import (
	"encoding/json"
	"github.com/vinizer4/go-expert-fullcycle/apis/internal/dto"
	"github.com/vinizer4/go-expert-fullcycle/apis/internal/entity"
	"github.com/vinizer4/go-expert-fullcycle/apis/internal/infra/database"
	"net/http"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDb database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDb,
	}
}

func (h *UserHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
