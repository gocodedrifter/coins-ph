package internal

import "errors"

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountAlreadyExist = errors.New("account already exist")
var ErrAccountCreation = errors.New("unable to create the account")
var ErrorCodeUnknown = errors.New("general error")
var ErrorInvalidArgument = errors.New("invalid argument")
