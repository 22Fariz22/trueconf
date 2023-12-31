package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/22Fariz22/trueconf/internal/user"
	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
	"github.com/go-chi/chi"
)

var l logger.Logger

type handler struct {
	uc user.UseCase
}

func NewHandler(uc user.UseCase) handler {
	return handler{uc: uc}
}

//SearchUsers вывод всех пользователей
func (h *handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	data, err := h.uc.SearchUsers(ctx)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(data.List)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if len(res) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

//GetUser получить данные пользователья по id
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "id")

	data, err := h.uc.GetUser(ctx, id)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	i, ok := data.List[id]
	if ok && !i.Deleted {
		res, err := json.Marshal(i)
		if err != nil {
			l.Errorf(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	if !ok && !i.Deleted || ok && i.Deleted {
		w.Write([]byte("Такого пользователя нету."))
	}
}

//CreateUser создать нового пользователя
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var newU entity.User

	if err := json.NewDecoder(r.Body).Decode(&newU); err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.uc.CreateUser(ctx, newU)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user created"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userNumber := chi.URLParam(r, "id")

	var updateUser entity.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.uc.UpdateUser(ctx, userNumber, updateUser)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//DeleteUser удалить пользователя по id
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := chi.URLParam(r, "id")

	err := h.uc.DeleteUser(ctx, id)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
