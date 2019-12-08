package res

import "errors"

type ErrorCode error

var (
	LoginErrorCode_Contains ErrorCode = errors.New("已存在该用户!")
)
