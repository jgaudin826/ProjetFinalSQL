package main

import (
	"ProjetFinalSQL/database"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ! The Department function is used to create the Department profile page. This page display all the information about a department
func Department(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/department/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/department.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	departmentUuid := strings.ReplaceAll(r.URL.Path, "/department/", "")
	if strings.Contains(departmentUuid, "/") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	department := database.GetDepartmentByUuid(departmentUuid, w, r)

	departmentPage := struct {
		Uuid         string
		Name         string
		Leader       string
		EmployeeList []database.EmployeeInfo
	}{
		Uuid:         department.Uuid,
		Name:         department.Name,
		Leader:       department.Department_leader_uuid,
		EmployeeList: database.GetEmployeesByDepartment(department.Uuid, w, r),
	}

	err = tmpl.Execute(w, departmentPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
