package main

import (
	"ProjetFinalSQL/database"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ! The Position function is used to create the position profile page. This page display all the information about a position
func Position(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/position/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/position.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	positionUuid := strings.ReplaceAll(r.URL.Path, "/position/", "")
	if strings.Contains(positionUuid, "/") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	position := database.GetPositionByUuid(positionUuid, w, r)

	departmentPage := struct {
		Uuid         string
		Name         string
		Salary       int
		EmployeeList []database.EmployeeInfo
	}{
		Uuid:         position.Uuid,
		Name:         position.Title,
		Salary:       position.Salary,
		EmployeeList: database.GetEmployeesByPosition(position.Uuid, w, r),
	}

	err = tmpl.Execute(w, departmentPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
