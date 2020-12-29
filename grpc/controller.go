package grpc

import (
	"context"
	user_service "go-grpc-samples/service"
)

type UserServiceGrpcServer struct {
	UnimplementedUserServiceServer
	userService user_service.IUserService
}

func (server UserServiceGrpcServer) GetUsers(context context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	ids := request.GetIds()

	users, err := server.userService.GetUsers(ids)

	if err != nil {
		return nil, err
	}

	usersResponse := make([]*User, len(users))

	for _, user := range users {
		usersResponse = append(usersResponse, &User{Id: user.Id, Name: user.Name})
	}

	return &GetUsersResponse{Users: usersResponse}, nil
}


func NewUserServiceGrpcServer(userService user_service.IUserService) UserServiceGrpcServer {
	return UserServiceGrpcServer{userService: userService}
}
