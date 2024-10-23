package main

import (
	"ProjetFinalSQL/database"
	"html/template"
	"net/http"
	"log"
)

type EmployeeInfo struct {
	Uuid          string
	First_name    string
	Last_name     string
	Email         string
	Phone_number  string
	Department_id string
	Position_id   string
	Superior_id   string
}

func Employee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employeeUuid := r.URL.Query().Get("uuid")

	log.Printf(employeeUuid)

	employee := database.GetEmployeeByUuid(employeeUuid, w, r)
	if (employee == database.Employee{}) {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("./templates/employee.html") 
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Employee EmployeeInfo
	}{
		Employee: EmployeeInfo{
			Uuid:          employee.Uuid, 
			First_name:    employee.First_name,
			Last_name:     employee.Last_name,
			Email:         employee.Email,
			Phone_number:  employee.Phone_number,
			Department_id: employee.Department_id,
			Position_id:   employee.Position_id,
			Superior_id:   employee.Superior_id,
		},
	})
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}