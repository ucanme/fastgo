package consts

import "errors"

const (
	PARAM_ERR_CODE =  100001
	SYSMSTEM_ERR_CODE = 100002

	USER_NOT_FOUND_CODE = 200001
	USER_PASS_INCORRECT_CODE = 200002
)

var PARAM_ERR = errors.New("param error")
var VALIDATE_ERR = errors.New("validate error")
var SYSMSTEM_ERR = errors.New("system error")

var USER_NOT_FOUND_ERR = errors.New("user not found")
var USER_PASS_INCORRECT_ERR = errors.New("password incorrect")