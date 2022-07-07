package routes

import (
	"github.com/foruscommunity/go-rest-api-example/pkg/controllers"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(controllers.CommonMiddleware)

	// Index (GET) (pkg/controllers/index.go)
	r.HandleFunc("/", controllers.IndexHandler).Methods("GET", "OPTIONS")

	// Users (GET, POST, PUT, DELETE) (pkg/controllers/users.go)
	r.HandleFunc("/users", controllers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.EditUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")

	return r
}
