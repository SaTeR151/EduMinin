package apperror

import "errors"

var ErrIdRequired = errors.New("id required")
var ErrDataNotFound = errors.New("data not found")
var ErrIncorrectRequestBody = errors.New("incorrect request body")

var ErrUserNotFound = errors.New("user not found")
var ErrUnauthorized = errors.New("unauthorized user")
var ErrGUIDRequired = errors.New("guid required")
var ErrTypecastJWT = errors.New("failed to typecast jwt claims")
var ErrIncorrectRefreshToken = errors.New("invalid refresh token format")

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserAlreadyAuthorization = errors.New("user already authorization")
var ErrUncorrectData = errors.New("uncorrect login or password")
