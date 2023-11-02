package api

import "net/http"

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	// 	GET http://localhost:3333/api/v1/users
	// Accept: application/json

	//вытащить данные из файла всех юзеров
	//передать в response данные из файла
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
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//GET http://localhost:3333/api/v1/users/0
	//Accept: application/json

	//получить номер  из request
	//достать из файла юзера по номеру
	//отдать юзера в response 
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// PATCH http://localhost:3333/api/v1/users/1
	// Content-Type: application/json

	// {
	//   "display_name": "TEST5"
	// }

	//получить номер из request и имя из json
	//поменять данные юзера в файле
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//DELETE http://localhost:3333/api/v1/users/2

	//получить номер из request
	//удалить его из файла(или поменять флажок deleted на True)
}
