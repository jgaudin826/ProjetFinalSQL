package main

import (
	"ProjetFinalSQL/functions"
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

// ! The main function is where the programme start. The function initialize all the page path. This is the root of the forums.
func main() {
	FileServer := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", FileServer))
	//! All pages redirection path
	http.HandleFunc("/", Home)
	http.HandleFunc("/team/", Team)
	http.HandleFunc("/employee/", Employee)
	http.HandleFunc("/department/", Department)
	http.HandleFunc("/position/", Position)
	// http.HandleFunc("/404", NotFound)

	//!Forms routes
	http.HandleFunc("/createEmployee", functions.CreateEmployee)
	http.HandleFunc("/updateEmployee", functions.UpdateEmployee)
	http.HandleFunc("/deleteEmployee", functions.DeleteEmployee)
	http.HandleFunc("/createTeam", functions.CreateTeam)
	http.HandleFunc("/updateTeam", functions.UpdateTeam)
	http.HandleFunc("/deleteTeam", functions.DeleteTeam)
	http.HandleFunc("/createDepartment", functions.CreateDepartment)
	http.HandleFunc("/updateDepartment", functions.UpdateDepartment)
	http.HandleFunc("/deleteDepartment", functions.DeleteDepartment)
	http.HandleFunc("/createPosition", functions.CreatePosition)
	http.HandleFunc("/updatePosition", functions.UpdatePosition)
	http.HandleFunc("/deletePosition", functions.DeletePosition)
	http.HandleFunc("/createLeave", functions.CreateLeave)
	http.HandleFunc("/updateLeave", functions.UpdateLeave)
	http.HandleFunc("/deleteLeave", functions.DeleteLeave)
	http.HandleFunc("/createEmployeeTeam", functions.CreateEmployeeTeam)
	http.HandleFunc("/deleteEmployeeTeam", functions.DeleteEmployeeTeam)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
