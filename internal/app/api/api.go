package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

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

func (h *handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	// data, err := openFile()
	// if err != nil {
	// 	l.Errorf(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

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

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser entity.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := openFile()
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data.Increment++
	newUser.CreatedAt = time.Now()
	incrStr := strconv.Itoa(data.Increment)

	data.List[incrStr] = newUser

	//маршалим data
	res, err := json.Marshal(data)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//записываем в файл
	err = os.WriteFile("users.json", res, 0666)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userNumber := chi.URLParam(r, "id")

	data, err := openFile()
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	i, ok := data.List[userNumber]
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
		w.Write([]byte("Такого юзера нету."))
	}
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userNumber := chi.URLParam(r, "id")

	var updateUser entity.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := openFile()
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//проверка на наличие такого юзера
	i, ok := data.List[userNumber]
	if ok && !i.Deleted {
		//копируем сущность User и меняем значение
		updUser := data.List[userNumber]
		updUser.DisplayName = updateUser.DisplayName
		data.List[userNumber] = updUser

		//маршалим data
		res, err := json.Marshal(data)
		if err != nil {
			l.Errorf(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//записываем в файл
		err = os.WriteFile("users.json", res, 0666)
		if err != nil {
			l.Errorf(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	deleteNumber := chi.URLParam(r, "id")

	data, err := openFile()
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//проверка на наличие такого юзера
	i, ok := data.List[deleteNumber]
	if ok && !i.Deleted {
		//копируем сущность User и меняем значение
		delUser := data.List[deleteNumber]
		delUser.Deleted = true
		data.List[deleteNumber] = delUser

		//маршалим data
		res, err := json.Marshal(data)
		if err != nil {
			l.Errorf(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//записываем в файл
		err = os.WriteFile("users.json", res, 0666)
		if err != nil {
			l.Errorf(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func openFile() (*entity.UserStore, error) {
	file, err := os.ReadFile("users.json")
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	data := entity.UserStore{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	return &data, nil
}
