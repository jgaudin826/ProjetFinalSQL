package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CheckErr(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

type Leave struct {
	Uuid         string
	EmployeeUuid string
	StartDate    time.Time
	EndDate      time.Time
	LeaveType    string
}

/*
!AddLeave function open data base and add leave to it with the INSERT INTO sql command she take as argument a Leave type and a writer and request.
*/
func AddLeave(leave Leave, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO leave (uuid, employee_uuid, start_date, end_date, leave_type) VALUES (?, ?, ?, ?, ?)")
	query.Exec(leave.Uuid, leave.EmployeeUuid, leave.StartDate, leave.EndDate, leave.LeaveType)
	CheckErr(err2, w, r)
	defer query.Close()
}

/*
!GetLeaveByUuid function is used to get a leave by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a Leave type.
*/
func GetLeaveByUuid(uuid string, w http.ResponseWriter, r *http.Request) Leave {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT * FROM leave WHERE uuid = ?", uuid)
	defer rows.Close()

	leave := Leave{}

	for rows.Next() {
		rows.Scan(&leave.Uuid, &leave.EmployeeUuid, &leave.StartDate, &leave.EndDate, &leave.LeaveType)
	}

	return leave
}

func GetLeaveByEmployeeUuid(employeeUuid string, w http.ResponseWriter, r *http.Request) []Leave {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM leave WHERE employee_uuid = ?", employeeUuid)
	defer rows.Close()

	var leaves []Leave

	for rows.Next() {
		leave := Leave{}
		rows.Scan(&leave.Uuid, &leave.EmployeeUuid, &leave.StartDate, &leave.EndDate, &leave.LeaveType)
		leaves = append(leaves, leave)
	}

	return leaves
}

/*
!UpdateLeaveInfo function is used to update leave information by using UPDATE sql command. She take as argument a Leave type, a writer, a request.
*/
func UpdateLeaveInfo(leave Leave, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE leave SET employee_uuid = ?, start_date = ?, end_date = ?, leave_type = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(leave.EmployeeUuid, leave.StartDate, leave.EndDate, leave.LeaveType, leave.Uuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 leave was affected")
	}
}

/*
!DeleteLeave function is used to delete a leave by using DELETE sql command. She take as argument an int, a writer, a request.
*/
func Deleteleave(leaveUuid string, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM leave WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(leaveUuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 leave was deleted")
	}
}
