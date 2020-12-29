package user_service

import (
	"errors"
	"go-grpc-samples/dbclient"
)

var ErrNotFound = errors.New("user not found")

type User struct {
	Id int64
	Name string
}

type IUserService interface {
	GetUser(id int64) (User, error)
	GetUsers(ids []int64) (map[int64]User, error)
}

type UserService struct {
	dbClient dbclient.BoltClient
}

func NewUserService(dbClient dbclient.BoltClient) UserService {
	return UserService{dbClient}
}

func (us *UserService) GetUser(id int64) (User, error) {
	return us.dbClient.GetUser(id)
}

func (us *UserService) GetUsers(ids []int64) (map[int64]User, error) {

	users := map[int64]User{}

	for _ , id := range ids {

		user, err := us.dbClient.GetUser(id)

		if err != nil {
			return users, err
		}

		users[id] = user
	}

	return users, nil
}
