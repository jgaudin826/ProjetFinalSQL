package main

import (
	"ProjetFinalSQL/database"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ! The Team function is used to create the Team profile page. This page display all the information about a team
func Team(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/team/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/team.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	teamUuid := strings.ReplaceAll(r.URL.Path, "/team/", "")
	if strings.Contains(teamUuid, "/") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	team := database.GetTeamByUuid(teamUuid, w, r)

	teamPage := struct {
		Uuid         string
		Name         string
		Leader       string
		EmployeeList []database.EmployeeInfo
	}{
		Uuid:         team.Uuid,
		Name:         team.Name,
		Leader:       team.Team_leader_name,
		EmployeeList: database.GetEmployeesByTeam(team.Uuid, w, r),
	}

	err = tmpl.Execute(w, teamPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
