package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code= 1000... user error
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	// code= 2000... art error
	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... cate error
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "Username already exist！",
	ERROR_PASSWORD_WRONG:   "Password Wrong!",
	ERROR_USER_NOT_EXIST:   "Username does not exist!",
	ERROR_TOKEN_EXIST:      "TOKEN does not exist,please login again!",
	ERROR_TOKEN_RUNTIME:    "TOKEN expired,please login again!",
	ERROR_TOKEN_WRONG:      "TOKEN error,please login again!",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN format error,please login again!",
	ERROR_USER_NO_RIGHT:    "User has no permissions!",

	ERROR_ART_NOT_EXIST: "Article does not exist!",

	ERROR_CATENAME_USED:  "This category already exists!",
	ERROR_CATE_NOT_EXIST: "This category dose not exists",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}