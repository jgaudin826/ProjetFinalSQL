package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
)

/*
! CreateTeam collects the user input from the corresponding form
! create a team in a database.team struct type
! sends it to the database function to store it
! redirects the user to the corresponding page
*/
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	teamName := r.FormValue("teamName")
	team := database.GetTeamByName(teamName, w, r)
	if (team != database.Team{}) {
		fmt.Println("Team already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Team+name+already+exist+!", http.StatusSeeOther)
		return
	}

	leaderFirstName := r.FormValue("leaderFirstName")
	leaderLastName := r.FormValue("leaderLastName")
	teamLeader := database.GetEmployeeByName(leaderFirstName, leaderLastName, w, r)
	if (teamLeader == database.Employee{}) {
		fmt.Println("Employee does not exist") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	teamUuid := GetNewUuid()

	newTeam := database.Team{Uuid: teamUuid, Team_leader_uuid: teamLeader.Uuid, Name: teamName}
	database.AddTeam(newTeam, w, r)

	http.Redirect(w, r, "/team/"+teamUuid+"?type=success&message=Team+created+successfuly+!", http.StatusSeeOther)
}

/*
! UpdateTeam collects the user input from the corresponding form
! create a team in a database.team struct type
! sends it to the database function to update it
! redirects the user to the corresponding page
*/
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	teamUuid := r.FormValue("teamUuid")
	team := database.GetTeamByUuid(teamUuid, w, r)
	if (team == database.Team{}) {
		fmt.Println("team does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/?type=error&message=Team+not+found+!", http.StatusSeeOther)
		return
	}

	newTeamName := r.FormValue("name")
	if newTeamName == "" {
		fmt.Println("name empty") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/team/"+teamUuid+"?type=error&message=Cannot+have+empty+name+!", http.StatusSeeOther)
		return
	}

	leaderFirstName := r.FormValue("leaderFirstName")
	leaderLastName := r.FormValue("leaderLastName")
	teamLeader := database.GetEmployeeByName(leaderFirstName, leaderLastName, w, r)
	if (teamLeader == database.Employee{}) {
		fmt.Println("Employee does not exist") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	newTeam := database.Team{Uuid: teamUuid, Team_leader_uuid: teamLeader.Uuid, Name: newTeamName}
	database.UpdateTeamInfo(newTeam, w, r)

	http.Redirect(w, r, "/team/"+teamUuid+"?type=success&message=Team+successfully+update+!", http.StatusSeeOther)
}

/*
! DeleteTeam collects the user input from the corresponding form
! sends it to the database function to delete it
! redirects the user to the corresponding page
*/
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	teamUuid := r.FormValue("teamUuid")
	team := database.GetTeamByUuid(teamUuid, w, r)
	if (team == database.Team{}) {
		fmt.Println("team does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/?type=error&message=Team+not+found+!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/team/"+teamUuid+"?type=error&message=Confirm+deletion+!", http.StatusSeeOther)
		return
	} else {
		database.DeleteTeam(teamUuid, w, r)
	}

	//Send confirmation message
	http.Redirect(w, r, "/?type=success&message=Team+deleted+!", http.StatusSeeOther)
}

/*
! CreateEmployeeTeam collects the user input from the corresponding form
! calls for the database function to add it
! redirects the user to the corresponding page
*/
func CreateEmployeeTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	teamUuid := r.FormValue("teamUuid")
	team := database.GetTeamByUuid(teamUuid, w, r)
	if (team == database.Team{}) {
		fmt.Println("team does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/search/?type=error&message=Team+not+found+!", http.StatusSeeOther)
		return
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	employee := database.GetEmployeeByName(firstName, lastName, w, r)
	if (employee == database.Employee{}) {
		fmt.Println("Employee does not exist") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	if database.ExistsEmployeeTeam(employee.Uuid, teamUuid, w, r) {
		fmt.Println("user already following this community")
		http.Redirect(w, r, "/team/"+teamUuid+"?type=error&message=Employee+was+already+part+of+this+team+!", http.StatusSeeOther)
	} else {
		database.AddEmployeeTeam(employee.Uuid, teamUuid, w, r)
		http.Redirect(w, r, "/team/"+teamUuid+"?type=success&message="+employee.First_name+"+"+employee.Last_name+"+is+now+part+of+team+"+team.Name+"+!", http.StatusSeeOther)
	}
}

/*
! DeleteEmployeeTeam collects the user input from the corresponding form
! calls for the database function to remove it
! redirects the user to the corresponding page
*/
func DeleteEmployeeTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	teamUuid := r.FormValue("teamUuid")
	team := database.GetTeamByUuid(teamUuid, w, r)
	if (team == database.Team{}) {
		fmt.Println("team does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/search/?type=error&message=Team+not+found+!", http.StatusSeeOther)
		return
	}

	employeeUuid := r.FormValue("employeeUuid")
	employee := database.GetEmployeeByUuid(employeeUuid, w, r)
	if (employee == database.Employee{}) {
		fmt.Println("Employee does not exist") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	if database.ExistsEmployeeTeam(employeeUuid, teamUuid, w, r) && employeeUuid != team.Team_leader_uuid {
		database.DeleteEmployeeTeam(employeeUuid, teamUuid, w, r)
		http.Redirect(w, r, "/team/"+teamUuid+"?type=success&message=Removed+"+employee.First_name+"+"+employee.Last_name+"+from+team+"+team.Name+"+!", http.StatusSeeOther)
	} else if employeeUuid == team.Team_leader_uuid {
		fmt.Println("employee cannot be removed because they are team leader")
		http.Redirect(w, r, "/team/"+teamUuid+"?type=error&message=Cannot+remove+team+leader+!+Plase+change+change+team+leader+before+removing+them+!", http.StatusSeeOther)
	} else {
		fmt.Println("employee already not part of this team")
		http.Redirect(w, r, "/team/"+teamUuid+"?type=error&message=Employee+was+already+not+part+of+this+team+!", http.StatusSeeOther)
	}
}
