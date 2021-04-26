package consts

import "errors"

const (
	PARAM_ERR_CODE =  100001
	SYSMSTEM_ERR_CODE = 100002

	USER_NOT_FOUND_CODE = 200001
	USER_PASS_INCORRECT_CODE = 200002
	USER_EXISTS_CODE = 200003

	DB_EXEC_ERR_CODE = 300001
	MAKE_APPOINT_FAIL_CODE = 300002
	TIME_ALREADY_OCCUPY_CODE = 300003
	GET_APPOINT_LIST_FAIL_CODE = 300004
	SIGN_IN_FAIL_CODE = 300005
	ADMIN_GET_APPOINT_FAIL_CODE = 400001


)

var PARAM_ERR = errors.New("param error")
var VALIDATE_ERR = errors.New("validate error")
var SYSMSTEM_ERR = errors.New("system error")

var USER_NOT_FOUND_ERR = errors.New("user not found")
var USER_PASS_INCORRECT_ERR = errors.New("password incorrect")
var USER_EXISTS_ERR = errors.New("user already exists")
var DB_EXEC_ERR = errors.New("db exec fail")
var MAKE_APPOINT_FAIL =  errors.New("make appoint fail")
