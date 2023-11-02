package entity

import "time"

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
		Deleted 		bool 			`json:"deleted"`
	}

	UserList  map[string]User

	UserStore struct {
		Increment int      `json:"increment"`
		List      UserList `json:"list"`
	}
)