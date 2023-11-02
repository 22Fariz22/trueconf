package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
	"github.com/go-chi/chi"
)

var l logger.Logger

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	data, err := openFile()
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := json.Marshal(data.List)
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if len(res) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// 	POST http://localhost:3333/api/v1/users
	// Content-Type: application/json

	// {
	//   "display_name": "TEST1",
	//   "email": "test"
	// }

	//создать юзера
	//вытащить имя и емейл из json
	//записать новые данные в файл и увеличить инкремент и также записать номер юзера

	var newUser entity.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		l.Errorf("error in Create(): ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("newUser: ", newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userNumber := chi.URLParam(r, "id")

	data, err := openFile()
	if err != nil {
		l.Errorf(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	i, ok := data.List[userNumber]
	if ok && !i.Deleted {
		res, err := json.Marshal(i)
		if err != nil {
			l.Errorf(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	if !ok && !i.Deleted || ok && i.Deleted {
		w.Write([]byte("Такого юзера нету."))
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// PATCH http://localhost:3333/api/v1/users/1
	// Content-Type: application/json

	// {
	//   "display_name": "TEST5"
	// }

	//получить номер из request и имя из json
	//поменять данные юзера в файле

	var updateUser entity.User

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		l.Errorf("error in Update(): ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("updateUser: ", updateUser.DisplayName)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//DELETE http://localhost:3333/api/v1/users/2

	//получить номер из request
	//удалить его из файла(или поменять флажок deleted на True)

	deleteNumber := chi.URLParam(r, "id")

	data, err := openFile()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	i, ok := data.List[deleteNumber]
	if ok && i.Deleted == false {
		i.Deleted = true
	}
}

func openFile() (*entity.UserStore, error) {
	file, err := os.ReadFile("users.json")
	if err != nil {
		l.Errorf("err readfile: ", err)
		return nil, err
	}

	data := entity.UserStore{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		l.Errorf("err unmarshall: ", err)
		return nil, err
	}

	return &data, nil
}
