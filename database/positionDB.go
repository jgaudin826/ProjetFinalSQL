package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Position struct {
	Uuid   string
	Title  string
	Salary int
}

/*
!AddPosition function open data base and add a position to it with the INSERT INTO sql command she take as argument a position type and a writer and request.
*/
func AddPosition(position Position, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO position (uuid, title, salary) VALUES (?, ?, ?)")
	query.Exec(position.Uuid, position.Title, position.Salary)
	CheckErr(err2, w, r)
	defer query.Close()
}

/*
!GetPositionByUuid function is used to get a position by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a Position type.
*/
func GetPositionByUuid(uuid string, w http.ResponseWriter, r *http.Request) Position {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM position WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	position := Position{}

	for rows.Next() {
		rows.Scan(&position.Uuid, &position.Title, &position.Salary)
	}

	return position
}

/*
!GetPositionByName function is used to get a position by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a position type.
*/
func GetPositionByName(positionName string, w http.ResponseWriter, r *http.Request) Position {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM position WHERE title = '" + positionName + "'")
	defer rows.Close()

	position := Position{}

	for rows.Next() {
		rows.Scan(&position.Uuid, &position.Title, &position.Salary)
	}

	return position
}

/*
!GetEmployeesByPosition function open data base and get users on a community by using the SELECT * FROM sql command she take as argument an int type and a writer and request and return a slice of User.
*/
func GetEmployeesByPosition(positionUuid string, w http.ResponseWriter, r *http.Request) []EmployeeInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT e.uuid, e.last_name, e.first_name, e.email, e.phone_number, e.department_uuid, d.name as department_name, e.position_uuid, p.title as position_name, e.superior_uuid, s.first_name || ' ' || s.last_name as superior_name FROM employee e LEFT JOIN department d ON e.department_uuid = d.uuid LEFT JOIN position p ON e.position_uuid = p.uuid LEFT JOIN employee s ON e.superior_uuid = s.uuid WHERE e.position_uuid='" + positionUuid + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	employeeList := make([]EmployeeInfo, 0)

	for rows.Next() {
		employee := EmployeeInfo{}
		err = rows.Scan(&employee.Uuid, &employee.Last_name, &employee.First_name, &employee.Email, &employee.Phone_number, &employee.Department_uuid, &employee.Department_name, &employee.Position_uuid, &employee.Position_name, &employee.Superior_uuid, &employee.Superior_name)
		CheckErr(err, w, r)

		employeeList = append(employeeList, employee)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return employeeList
}

/*
!UpdatePositionInfo function is used to update position information by using UPDATE sql command. She take as argument a position type, a writer, a request.
*/
func UpdatePositionInfo(position Position, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE position SET title= ?, salary = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(position.Title, position.Salary, position.Uuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 position was affected")
	}
}

/*
!DeletePosition function is used to delete a position by using DELETE sql command. She take as argument an int, a writer, a request.
*/
func DeletePosition(positionUuid string, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM position WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(positionUuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 position was deleted")
	}
}
