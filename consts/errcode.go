package consts

import "errors"

const (
	PARAM_ERR_CODE =  100001
	ACCOUTN_NOT_LOGIN =  100002




	DB_EXEC_ERR_CODE = 300001
	REQUEST_FAIL_CODE = 400001

	DB_ROWS_AFFECTED_ZERO_CODE= 300002
	DB_QUERY_ERR_CODE = 300001

	DB_QUERY_NOT_EXIST_CODE = 300003


)

var PARAM_ERR = errors.New("param error")
var VALIDATE_ERR = errors.New("validate error")
var SYSMSTEM_ERR = errors.New("system error")

var USER_NOT_FOUND_ERR = errors.New("user not found")
var USER_PASS_INCORRECT_ERR = errors.New("password incorrect")
var USER_EXISTS_ERR = errors.New("user already exists")
var DB_EXEC_ERR = errors.New("db exec fail")
var DB_ROWS_AFFECTED_ZERO_ERR = errors.New("db exec rows affected zero")
var DB_QUERY_FAIL = errors.New("db query fail")
var MAKE_APPOINT_FAIL =  errors.New("make appoint fail")
var REQUEST_FAIL = errors.New("request_api_fail")
var DB_QUERY_NOT_EXIST_ERR = errors.New("db query not exist")
