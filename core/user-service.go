package core

import (
	"go-grpc-samples/dbclient"
	"go-grpc-samples/user-service"
)



type UserService struct {
	dbClient dbclient.BoltClient
}

func NewUserService(dbClient dbclient.BoltClient) UserService {
	return UserService{dbClient}
}

func (us UserService) GetUser(id int64) (user_service.User, error) {
	return us.dbClient.GetUser(id)
}

func (us UserService) GetUsers(ids []int64) (map[int64]user_service.User, error) {

	users := map[int64]user_service.User{}

	for _ , id := range ids {

		user, err := us.dbClient.GetUser(id)

		if err != nil {
			return users, err
		}

		users[id] = user
	}

	return users, nil
}
