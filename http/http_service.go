package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-grpc-samples/core"
	"go-grpc-samples/dbclient"
	"io/ioutil"
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
	var users []int
	json.Unmarshal(content, &users)



}

func (s server) Start(address string) {
	s.router.HandleFunc("/users", s.HandleUsers)
	http.Handle("/", s.router)
}

func NewServer() server {
	return server {
		userService: core.NewUserService(dbclient.GetDatabase()),
		router: mux.NewRouter(),
	}
}
