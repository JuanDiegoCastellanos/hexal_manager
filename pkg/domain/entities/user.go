package entities

type User struct {
	ID       int64  `json:"id"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
}
