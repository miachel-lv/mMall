package network

type UserService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (self *UserService) Verify() bool {
	//TODO
	return false
}