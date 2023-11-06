package filejson

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
)

var l logger.Logger

// inFileRepo структура для стоража инфайл
type inFileRepo struct {
	fileName string
}

func NewRepo(fileName string) *inFileRepo {
	return &inFileRepo{fileName: fileName}
}

func (f *inFileRepo) CreateUser(ctx context.Context, newU entity.User) error {
	data := entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf(err)
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf(err)
		return err
	}

	data.Increment++
	incrStr := strconv.Itoa(data.Increment)
	data.List[incrStr] = newU

	res, err := json.Marshal(data)
	if err != nil {
		l.Errorf(err)
		return err
	}

	//записываем в файл
	err = os.WriteFile("users.json", res, 0666)
	if err != nil {
		l.Errorf(err)
		return err
	}

	return nil
}

func (f *inFileRepo) DeleteUser(ctx context.Context, id string) error {
	data := entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf(err)
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf(err)
		return err
	}

	//проверка на наличие такого юзера
	i, ok := data.List[id]
	if ok && !i.Deleted {
		//копируем сущность User и меняем значение
		delUser := data.List[id]
		delUser.Deleted = true
		data.List[id] = delUser

		//маршалим data
		res, err := json.Marshal(data)
		if err != nil {
			l.Errorf(err)
			return err
		}

		//записываем в файл
		err = os.WriteFile("users.json", res, 0666)
		if err != nil {
			l.Errorf(err)
			return err
		}
	}

	return nil
}

func (f *inFileRepo) GetUser(ctx context.Context, id string) (*entity.UserStore, error) {
	fmt.Println("Repo GetUser()")
	data := entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	return &data, nil
}

func (f *inFileRepo) UpdateUser(ctx context.Context, id string, updateUser entity.User) error {
	data := entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf(err)
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf(err)
		return err
	}

	//проверка на наличие такого юзера
	i, ok := data.List[id]
	if ok && !i.Deleted {
		//копируем сущность User и меняем значение
		updUser := data.List[id]
		updUser.DisplayName = updateUser.DisplayName
		data.List[id] = updUser

		//маршалим data
		res, err := json.Marshal(data)
		if err != nil {
			l.Errorf(err)
			return err
		}

		//записываем в файл
		err = os.WriteFile("users.json", res, 0666)
		if err != nil {
			l.Errorf(err)
			return err
		}
	}

	return nil
}

func (f *inFileRepo) SearchUsers(ctx context.Context) (*entity.UserStore, error) {
	data := entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	return &data, nil
}
