package apperror

import "errors"

var ErrIdRequired = errors.New("id required")
var ErrDataNotFound = errors.New("data not found")
var ErrIncorrectRequestBody = errors.New("incorrect request body")
