package model

type User struct {
	UserId    int64
	Username  string
	AvatarURL string
	Password  string
	Email     string
	Role      int64
}
