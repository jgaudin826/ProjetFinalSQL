package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	Uuid         string
	Last_name    string
	First_name   string
	Email        string
	Phone_number string
	Department_id string
	Position_id  string
	Superior_id  string
}

// AddEmployee opens the database connection and adds an employee to it using the INSERT INTO SQL command.
// It takes an Employee struct, http.ResponseWriter, and *http.Request as arguments.
func AddEmployee(employee Employee, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO employee (uuid, Last_name, First_name, Email, Phone_number, Department_id, Position_id, Superior_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	query.Exec(employee.Uuid, employee.Last_name, employee.First_name, employee.Email, employee.Phone_number, employee.Department_id, employee.Position_id, employee.Superior_id)
	CheckErr(err2, w, r)
	defer query.Close()
}

// GetEmployeeByUuid retrieves an employee by their UUID using the SELECT * FROM SQL command.
// It takes a UUID string, http.ResponseWriter, and *http.Request as arguments, and returns an Employee struct.
func GetEmployeeByUuid(uuid string, w http.ResponseWriter, r *http.Request) Employee {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM employee WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	employee := Employee{}

	for rows.Next() {
		rows.Scan(&employee.Uuid, &employee.Last_name, &employee.First_name, &employee.Email, &employee.Phone_number, &employee.Department_id, &employee.Position_id, &employee.Superior_id)
	}

	return employee
}

// UpdateEmployeeInfo updates an employee's information using the UPDATE SQL command.
// It takes an Employee struct, http.ResponseWriter, and *http.Request as arguments.
func UpdateEmployeeInfo(employee Employee, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("UPDATE employee SET Last_name = ?, First_name = ?, Email = ?, Phone_number = ?, Department_id = ?, Position_id = ?, Superior_id = ? WHERE Uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(employee.Last_name, employee.First_name, employee.Email, employee.Phone_number, employee.Department_id, employee.Position_id, employee.Superior_id, employee.Uuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error: More than 1 employee was affected")
	}
}

// DeleteEmployee deletes an employee from the database using the DELETE SQL command.
// It takes an employee UUID string, http.ResponseWriter, and *http.Request as arguments.
func DeleteEmployee(employeeUuid string, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("DELETE FROM employee WHERE Uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(employeeUuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error: More than 1 employee was deleted")
	}
}
