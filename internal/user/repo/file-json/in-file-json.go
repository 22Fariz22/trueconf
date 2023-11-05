package filejson

import (
	"context"
	"encoding/json"
	"os"

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

func (f *inFileRepo) CreateUser(ctx context.Context) {}
func (f *inFileRepo) DeleteUser(ctx context.Context) {}
func (f *inFileRepo) GetUser(ctx context.Context)    {}
func (f *inFileRepo) UpdateUser(ctx context.Context) {}

func (f *inFileRepo) SearchUsers(ctx context.Context) (*entity.UserStore, error) {
	file, err := os.ReadFile(f.fileName)
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
