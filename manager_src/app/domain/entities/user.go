package entities

type User struct {
	ID        int64  `json:"id"`
	Nick_name string `json:"nick_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleId    int64  `json:"role_id"`
}

func NewUser(id int64, nickName string, email string, password string, roleId int64) *User {
	return &User{
		ID:        id,
		Nick_name: nickName,
		Email:     email,
		Password:  password,
		RoleId:    roleId,
	}
}
