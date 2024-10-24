package main

import (
	"html/template"
	"log"
	"net/http"
	"ProjetFinalSQL/database"
)

type DepartmentInfo struct {
	Uuid string
	Name string
	Department_leader_uuid string
	Department_leader_name string
}

type TeamInfo struct {
	Uuid string
	Name string
}

// !The Home function is used to create the home page which is the main page of the forums. She take as arguments a writer and a request.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employeesInfo := database.GetAllEmployees(w, r) 
	departmentsInfo := database.GetAllDepartments(w, r)
	positionsInfo := database.GetAllPositions(w, r)
	teamsInfo := database.GetAllTeams(w, r)

	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Employees   []database.EmployeeInfo 
		Departments []database.DepartmentInfo 
		Positions   []database.Position  
		Teams       []database.TeamInfo      
	}{
		Employees:   employeesInfo,
		Departments: departmentsInfo,
		Positions:   positionsInfo,
		Teams:       teamsInfo,
	})
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
