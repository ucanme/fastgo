package consts

import "errors"

const (
	PARAM_ERR_CODE =  100001
)

var PARAM_ERR = errors.New("param error")
var VALIDATE_ERR = errors.New("validate error")