package e

const (
	SUCCESS       = 200
	InvalidParams = 400
	ERROR         = 500

	//管理员错误
	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuthToken                 = 30003
	ErrorAuth                      = 30004

	//对象存储错误
	ErrorOss        = 50001
	ErrorUploadFile = 50002
)