package api

import (
	"encoding/json"
	"net/http"

	"github.com/22Fariz22/trueconf/internal/user"
	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
	"github.com/go-chi/chi"
)

type handler struct {
	l  logger.Logger
	uc user.UseCase
}

func NewHandler(l logger.Logger, uc user.UseCase) handler {
	return handler{l: l, uc: uc}
}

// SearchUsers вывод всех пользователей
func (h *handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.uc.SearchUsers(ctx, h.l)
	if err != nil {
		h.l.Errorf("err in SearchUsers()->h.uc.SearchUsers(): ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(data.List)
	if err != nil {
		h.l.Errorf("err in SearchUsers()->json.Marshal: ", err)
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

// GetUser получить данные пользователья по id
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	data, err := h.uc.GetUser(ctx, h.l, id)
	if err != nil {
		h.l.Errorf("err %w in GetUser()->h.uc.GetUser() with id:%s", err, id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	i, ok := data.List[id]
	if ok && !i.Deleted {
		res, err := json.Marshal(i)
		if err != nil {
			h.l.Errorf("err %w in GetUser()->json.Marshal() with i:%v", err, i)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	if !ok && !i.Deleted || ok && i.Deleted {
		w.WriteHeader(http.StatusNotFound)
	}
}

// CreateUser создать нового пользователя
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newU *entity.User

	if err := json.NewDecoder(r.Body).Decode(&newU); err != nil {
		h.l.Errorf("err %w in CreateUser()->json.NewDecoder()", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.uc.CreateUser(ctx, h.l, newU)
	if err != nil {
		h.l.Errorf("err %w in CreateUser()->h.uc.CreateUser() with newU.DispalyName:%s", err, newU.DisplayName)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(struct{}{})
	w.Write(res)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userNumber := chi.URLParam(r, "id")

	var updateUser *entity.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		h.l.Errorf("err %w in UpdateUser()->json.NewDecoder()", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.uc.UpdateUser(ctx, h.l, userNumber, updateUser)
	if err != nil {
		h.l.Errorf("err %w in UpdateUser()->h.uc.UpdateUser() with userNumber %s", err, userNumber)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser удалить пользователя по id
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	err := h.uc.DeleteUser(ctx, h.l, id)
	if err != nil {
		h.l.Errorf("err %w in DeleteUser()->h.uc.DeleteUser() with id %s", err, id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
