package network

type RegisterService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (self *RegisterService) Verify() bool {
	//TODO
	return false
}