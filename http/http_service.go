package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-grpc-samples/core"
	"go-grpc-samples/dbclient"
	"go-grpc-samples/service"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	userService core.UserService
	router *mux.Router
}

func (s server) HandleUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var ids []int64
	json.Unmarshal(content, &ids)

	users, err := s.userService.GetUsers(ids)

	var usersResponse []service.User

	for _, user := range users {
		usersResponse = append(usersResponse, service.User { Name: user.Name, Id: user.Id})
	}

	bytes, _ := json.Marshal(usersResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

func (s server) Start(address string) {

	s.router.HandleFunc("/users", s.HandleUsers)
	http.Handle("/", s.router)

	fmt.Println("Starting HTTP service on address", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal("Error starting http server on port", address)
	}
}

func NewServer(db dbclient.BoltClient) server {


	return server {
		userService: core.NewUserService(db),
		router: mux.NewRouter(),
	}
}
