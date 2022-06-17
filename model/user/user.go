package user

type User struct {
	Id        int
	Username  string
	Password  string
	Salt      string
	Email     string
	State     int
	Avatar    string
	Role      string
	Signature string
}

type UserToken struct {
	Id       int
	Username string
	Email    string
}

func (u *User) TableName() string {
	return "user"
}
