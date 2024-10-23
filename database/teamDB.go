package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Team struct {
	Uuid             string
	Team_leader_uuid string
	Name             string
}

/*
!AddTeam function open data base and add team to it with the INSERT INTO sql command she take as argument a team type and a writer and request.
*/
func AddTeam(team Team, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO team (uuid, team_leader_uuid, name) VALUES (?, ?, ?)")
	query.Exec(team.Uuid, team.Team_leader_uuid, team.Name)
	CheckErr(err2, w, r)
	defer query.Close()
}

/*
!GetTeamByUuid function is used to get a team by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a team type.
*/
func GetTeamByUuid(uuid string, w http.ResponseWriter, r *http.Request) Team {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM team WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	team := Team{}

	for rows.Next() {
		rows.Scan(&team.Uuid, &team.Team_leader_uuid, &team.Name)
	}

	return team
}

/*
!GetTeamByName function is used to get a team by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a team type.
*/
func GetTeamByName(teamName string, w http.ResponseWriter, r *http.Request) Team {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM team WHERE name = '" + teamName + "'")
	defer rows.Close()

	team := Team{}

	for rows.Next() {
		rows.Scan(&team.Uuid, &team.Team_leader_uuid, &team.Name)
	}

	return team
}

/*
!UpdateTeamInfo function is used to update team information by using UPDATE sql command. She take as argument a team type, a writer, a request.
*/
func UpdateTeamInfo(team Team, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE team SET team_leader_uuid = ?, name = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(team.Team_leader_uuid, team.Name, team.Uuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 team was affected")
	}
}

/*
!DeleteTeam function is used to delete a team by using DELETE sql command. She take as argument an int, a writer, a request.
*/
func DeleteTeam(teamUuid string, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM team WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(teamUuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 team was deleted")
	}
}

/*
!AddEmployeeTeam function open data base and add team to it with the INSERT INTO sql command she take as argument a team type and a writer and request.
*/
func AddEmployeeTeam(employeeUuid string, teamUuid string, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO employee_team (employee_uuid, team_uuid) VALUES (?, ?)")
	query.Exec(employeeUuid, teamUuid)
	CheckErr(err2, w, r)
	defer query.Close()
}

/*
!ExistsEmployeeTeam function open data base and check if a employee is in a team by using the SELECT * FROM sql command she take as argument two int type and a writer and request and return a boolean.
*/
func ExistsEmployeeTeam(employeeUuid string, teamUuid string, w http.ResponseWriter, r *http.Request) bool {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM employee_team WHERE employee_uuid = '" + employeeUuid + "' AND team_uuid = '" + teamUuid + "'")
	defer rows.Close()

	type EmployeeTeam struct {
		EmployeeUuid string
		TeamUuid     string
	}
	employee_team := EmployeeTeam{}

	for rows.Next() {
		rows.Scan(&employee_team.EmployeeUuid, &employee_team.TeamUuid)
	}

	return employee_team != EmployeeTeam{}
}

/*
!DeleteEmployeeTeam function is used to delete a team by using DELETE sql command. She take as argument an int, a writer, a request.
*/
func DeleteEmployeeTeam(employeeUuid string, teamUuid string, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM employee_team WHERE employee_uuid = ? AND team_uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(employeeUuid, teamUuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 team was deleted")
	}
}
