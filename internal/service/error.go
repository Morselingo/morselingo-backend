package service

import "errors"

var ErrorUserAlreadyExists = errors.New("user already exists")
var ErrorCreateUserFailed = errors.New("failed to create user")
var ErrorUserNotFound = errors.New("user not found")

var ErrorAuthenticationFailed = errors.New("invalid credentials")
var ErrorFailedToHashPassword = errors.New("failed to hash password")
