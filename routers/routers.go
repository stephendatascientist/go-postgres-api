package routers

import (
	"github.com/gorilla/mux"
	"github.com/stephendatascientist/go-postgres-api/controllers"
	"gorm.io/gorm"
)

func Routers(db *gorm.DB) *mux.Router {

	// Create repository
	repository := &controllers.EmployeeRepository{DB: db}

	// Create router
	r := mux.NewRouter()

	// Configure path
	r.HandleFunc("/employee", repository.GetEmployees).Methods("GET")
	r.HandleFunc("/employee", repository.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", repository.GetEmployee).Methods("GET")
	r.HandleFunc("/employee/{id}", repository.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", repository.DeleteEmployee).Methods("DELETE")

	return r
}
