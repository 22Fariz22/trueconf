package filejson

import (
	"bufio"
	"context"
	"io"
)

//inFileRepository структура для стоража инфайл
type inFileRepository struct {
	file          io.ReadWriteCloser
	// memoryStorage storage.MemoryStorage
	reader        *bufio.Reader
}

func NewInFileRepository(){

}

func(f *inFileRepository)CreateUser(ctx context.Context){}
func(f *inFileRepository)DeleteUser(ctx context.Context){}
func(f *inFileRepository)GetUser(ctx context.Context){}
func(f *inFileRepository)UpdateUser(ctx context.Context){}
func(f *inFileRepository)SearchUser(ctx context.Context){}