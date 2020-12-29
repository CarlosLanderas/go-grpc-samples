package service

import "errors"

var ErrNotFound = errors.New("user not found")

type User struct {
	Id int64
	Name string
}

type IUserService interface {
	GetUser(id int64) (User, error)
	GetUsers(ids []int64) (map[int64]User, error)
}
