package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Department struct {
	Uuid                 string
	Name                 string
	Department_leader_id string
}

// AddDepartment opens the database connection and adds a department to it using the INSERT INTO SQL command.
// It takes a Department struct, http.ResponseWriter, and *http.Request as arguments.
func AddDepartment(department Department, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO department (uuid, name, department_leader_id) VALUES (?, ?, ?)")
	query.Exec(department.Uuid, department.Name, department.Department_leader_id)
	CheckErr(err2, w, r)
	defer query.Close()
}

// GetDepartmentByUuid retrieves a department by its UUID using the SELECT * FROM SQL command.
// It takes a UUID string, http.ResponseWriter, and *http.Request as arguments, and returns a Department struct.
func GetDepartmentByUuid(uuid string, w http.ResponseWriter, r *http.Request) Department {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM department WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	department := Department{}

	for rows.Next() {
		rows.Scan(&department.Uuid, &department.Name, &department.Department_leader_id)
	}

	return department
}

// UpdateDepartmentInfo updates a department's information using the UPDATE SQL command.
// It takes a Department struct, http.ResponseWriter, and *http.Request as arguments.
func UpdateDepartmentInfo(department Department, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("UPDATE department SET name = ?, department_leader_id = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(department.Name, department.Department_leader_id, department.Uuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error: More than 1 department was affected")
	}
}

// DeleteDepartment deletes a department from the database using the DELETE SQL command.
// It takes a department UUID string, http.ResponseWriter, and *http.Request as arguments.
func DeleteDepartment(departmentUuid string, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("DELETE FROM department WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(departmentUuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error: More than 1 department was deleted")
	}
}
