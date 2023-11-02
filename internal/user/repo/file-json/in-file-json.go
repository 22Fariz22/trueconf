package filejson

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/22Fariz22/trueconf/internal/config"
	"github.com/22Fariz22/trueconf/internal/user/entity"
)

// inFileRepository структура для стоража инфайл
type inFileRepository struct {
	file   io.ReadWriteCloser
	reader *bufio.Reader
}

// Consumer структура консьюмера
type Consumer struct {
	File   *os.File
	reader *bufio.Reader
}

// NewConsumer  создание консьюмера
func NewConsumer(cfg config.Config) (*Consumer, error) {
	file, err := os.OpenFile(cfg.Filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		File:   file,
		reader: bufio.NewReader(file),
	}, nil
}

// New инициализация консьюмера
func New(cfg *config.Config) usecase.Repository {
	st := storage.New()

	consumer, err := NewConsumer(*cfg)
	if err != nil {
		log.Println(err)
	}

	return &inFileRepository{
		file:   consumer.File,
		reader: consumer.reader,
	}
}

// Init инициализация консьмера
func (f *inFileRepository) Init() error {
	scanner := bufio.NewScanner(f.file)

	for scanner.Scan() {
		txt := scanner.Text()
		var u entity.User
		err := json.Unmarshal([]byte(txt), &u)
		if err != nil {
			return err
		}
		f.memoryStorage.Insert()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return nil
}

func (f *inFileRepository) CreateUser(ctx context.Context) {}
func (f *inFileRepository) DeleteUser(ctx context.Context) {}
func (f *inFileRepository) GetUser(ctx context.Context)    {}
func (f *inFileRepository) UpdateUser(ctx context.Context) {}
func (f *inFileRepository) SearchUser(ctx context.Context) {}
