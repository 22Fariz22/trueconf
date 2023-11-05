package filejson

import (
	"context"
	"encoding/json"
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

func (f *inFileRepo) DeleteUser(ctx context.Context) {

}

func (f *inFileRepo) GetUser(ctx context.Context, id int) (*entity.UserStore, error) {
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

func (f *inFileRepo) UpdateUser(ctx context.Context) {}

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
