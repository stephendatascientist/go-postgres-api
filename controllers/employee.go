package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stephendatascientist/go-postgres-api/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func (er *EmployeeRepository) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")

	// Check empty request body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var employee models.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)

	result := er.DB.Create(&employee)

	fmt.Println(result.RowsAffected)

	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(employee)
}

func (er *EmployeeRepository) GetEmployees(w http.ResponseWriter, r *http.Request) {

	var employees []models.Employee
	er.DB.Find(&employees)

	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(employees)
}

func (er *EmployeeRepository) GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])

	var employee models.Employee
	er.DB.First(&employee, params["id"])

	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(employee)
}

func (er *EmployeeRepository) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])

	var employee models.Employee
	er.DB.First(&employee, params["id"])

	// Check empty request body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	_ = json.NewDecoder(r.Body).Decode(&employee)
	er.DB.Save(&employee)

	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(employee)
}

func (er *EmployeeRepository) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])

	var employee models.Employee
	err := er.DB.Delete(&employee, params["id"])
	if err == nil {
		json.NewEncoder(w).Encode("Could not delete an employee")
	}
}
