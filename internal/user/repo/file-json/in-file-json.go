package filejson

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
)

// inFileRepo структура для стоража инфайл
type inFileRepository struct {
	fileName string
}

func NewRepo(fileName string) *inFileRepository {
	return &inFileRepository{fileName: fileName}
}

func (f *inFileRepository) CreateUser(ctx context.Context, l logger.Logger, newU *entity.User) error {
	data := &entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf("err %w in repository CreateUser()->os.ReadFile()", err)
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf("err %w in repository CreateUser()->json.Unmarshal()", err)
		return err
	}

	data.Increment++
	incrStr := strconv.Itoa(data.Increment)
	data.List[incrStr] = *newU

	res, err := json.Marshal(data)
	if err != nil {
		l.Errorf("err %w in repository CreateUser()->json.Marshal() with increment: %v", err, data.Increment)
		return err
	}

	//записываем в файл
	err = os.WriteFile("users.json", res, 0666)
	if err != nil {
		l.Errorf("err %w in repository CreateUser()->os.WriteFile()", err)
		return err
	}

	return nil
}

func (f *inFileRepository) DeleteUser(ctx context.Context, l logger.Logger, id string) error {
	data := entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf("err %w in repository DeleteUser()->os.ReadFile()", err)
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf("err %w in repository DeleteUser()->json.Unmarshal()", err)
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
			l.Errorf("err %w in repository DeleteUser()->json.Marshal()", err)
			return err
		}

		//записываем в файл
		err = os.WriteFile("users.json", res, 0666)
		if err != nil {
			l.Errorf("err %w in repository DeleteUser()->os.WriteFile()", err)
			return err
		}
	}

	return nil
}

func (f *inFileRepository) GetUser(ctx context.Context, l logger.Logger, id string) (*entity.UserStore, error) {
	data := &entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf("err %w in repository GetUser()->os.ReadFile()", err)
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf("err %w in repository GetUser()->json.Unmarshal()", err)
		return nil, err
	}

	return data, nil
}

func (f *inFileRepository) UpdateUser(ctx context.Context, l logger.Logger, id string, updateUser *entity.User) error {
	data := &entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf("err %w in repository UpdateUser()->os.ReadFile()", err)
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf("err %w in repository UpdateUser()->json.Unmarshal()", err)
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
			l.Errorf("err %w in repository UpdateUser()->json.Marshal()", err)
			return err
		}

		//записываем в файл
		err = os.WriteFile("users.json", res, 0666)
		if err != nil {
			l.Errorf("err %w in repository UpdateUser()->os.WriteFile()", err)
			return err
		}
	}

	return nil
}

func (f *inFileRepository) SearchUsers(ctx context.Context, l logger.Logger) (*entity.UserStore, error) {
	data := &entity.UserStore{}

	file, err := os.ReadFile(f.fileName)
	if err != nil {
		l.Errorf("err %w in repository SearchUsers()->os.ReadFile()", err)
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		l.Errorf("err %w in repository SearchUsers()->json.Unmarshal()", err)
		return nil, err
	}

	return data, nil
}
