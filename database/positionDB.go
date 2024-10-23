package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Position struct {
	Uuid   string
	Name   string
	Salary int
}

/*
!AddPosition function open data base and add a position to it with the INSERT INTO sql command she take as argument a position type and a writer and request.
*/
func AddPosition(position Position, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO position (uuid, name, salary) VALUES (?, ?, ?)")
	query.Exec(position.Uuid, position.Name, position.Salary)
	CheckErr(err2, w, r)
	defer query.Close()
}

/*
!GetPositionByUuid function is used to get a position by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a Position type.
*/
func GetPositionByUuid(uuid string, w http.ResponseWriter, r *http.Request) Position {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM position WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	position := Position{}

	for rows.Next() {
		rows.Scan(&position.Uuid, &position.Name, &position.Salary)
	}

	return position
}

/*
!GetPositionByName function is used to get a position by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a position type.
*/
func GetPositionByName(positionName string, w http.ResponseWriter, r *http.Request) Position {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM position WHERE name = '" + positionName + "'")
	defer rows.Close()

	position := Position{}

	for rows.Next() {
		rows.Scan(&position.Uuid, &position.Name, &position.Salary)
	}

	return position
}

/*
!UpdatePositionInfo function is used to update position information by using UPDATE sql command. She take as argument a position type, a writer, a request.
*/
func UpdatePositionInfo(position Position, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE position SET name= ?, salary = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(position.Name, position.Salary, position.Uuid)
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
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
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
