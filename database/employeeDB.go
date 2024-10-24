package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	Uuid            string
	Last_name       string
	First_name      string
	Email           string
	Phone_number    string
	Department_uuid string
	Position_uuid   string
	Superior_uuid   string
}

type EmployeeInfo struct {
	Uuid            string
	Last_name       string
	First_name      string
	Email           string
	Phone_number    string
	Department_uuid string
	Department_name string
	Position_uuid   string
	Position_name   string
	Superior_uuid   string
	Superior_name   string
}

// AddEmployee opens the database connection and adds an employee to it using the INSERT INTO SQL command.
// It takes an Employee struct, http.ResponseWriter, and *http.Request as arguments.
func AddEmployee(employee Employee, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO employee (uuid, last_name, first_name, email, phone_number, department_uuid, position_uuid, superior_uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	query.Exec(employee.Uuid, employee.Last_name, employee.First_name, employee.Email, employee.Phone_number, employee.Department_uuid, employee.Position_uuid, employee.Superior_uuid)
	CheckErr(err2, w, r)
	defer query.Close()
}

// GetEmployeeByUuid retrieves an employee by their UUID using the SELECT * FROM SQL command.
// It takes a UUID string, http.ResponseWriter, and *http.Request as arguments, and returns an Employee struct.
func GetEmployeeByUuid(uuid string, w http.ResponseWriter, r *http.Request) EmployeeInfo {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, _ := db.Query("SELECT e.uuid, e.last_name, e.first_name, e.email, e.phone_number, e.department_uuid, d.name as department_name, e.position_uuid, p.title as position_name, e.superior_uuid, s.first_name || ' ' || s.last_name as superior_name FROM employee e LEFT JOIN department d ON e.department_uuid = d.uuid LEFT JOIN position p ON e.position_uuid = p.uuid LEFT JOIN employee s ON e.superior_uuid = s.uuid WHERE e.uuid = '" + uuid + "'")
	defer rows.Close()

	employee := EmployeeInfo{}

	for rows.Next() {
		rows.Scan(&employee.Uuid, &employee.Last_name, &employee.First_name, &employee.Email, &employee.Phone_number, &employee.Department_uuid, &employee.Department_name, &employee.Position_uuid, &employee.Position_name, &employee.Superior_uuid, &employee.Superior_name)
	}

	return employee
}

// GetEmployeeByName retrieves an employee by their first and last name using the SELECT * FROM SQL command.
// It takes a firstname and lastname string, http.ResponseWriter, and *http.Request as arguments, and returns an Employee struct.
func GetEmployeeByName(firstName string, lastName string, w http.ResponseWriter, r *http.Request) Employee {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM employee WHERE last_name = '" + lastName + "' AND first_name = '" + firstName + "'")
	defer rows.Close()

	employee := Employee{}

	for rows.Next() {
		rows.Scan(&employee.Uuid, &employee.Last_name, &employee.First_name, &employee.Email, &employee.Phone_number, &employee.Department_uuid, &employee.Position_uuid, &employee.Superior_uuid)
	}

	return employee
}

// UpdateEmployeeInfo updates an employee's information using the UPDATE SQL command.
// It takes an Employee struct, http.ResponseWriter, and *http.Request as arguments.
func UpdateEmployeeInfo(employee Employee, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("UPDATE employee SET last_name = ?, first_name = ?, email = ?, phone_number = ?, department_uuid = ?, position_uuid = ?, superior_uuid = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(employee.Last_name, employee.First_name, employee.Email, employee.Phone_number, employee.Department_uuid, employee.Position_uuid, employee.Superior_uuid, employee.Uuid)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error: More than 1 employee was affected")
	}
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) []EmployeeInfo {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, _ := db.Query("SELECT e.uuid, e.last_name, e.first_name, e.email, e.phone_number, e.department_uuid, d.name as department_name, e.position_uuid, p.title as position_name, e.superior_uuid, s.first_name || ' ' || s.last_name as superior_name FROM employee e LEFT JOIN department d ON e.department_uuid = d.uuid LEFT JOIN position p ON e.position_uuid = p.uuid LEFT JOIN employee s ON e.superior_uuid = s.uuid")
	defer rows.Close()

	var employees []EmployeeInfo

	for rows.Next() {
		employee := EmployeeInfo{}
		rows.Scan(&employee.Uuid, &employee.Last_name, &employee.First_name, &employee.Email, &employee.Phone_number, &employee.Department_uuid, &employee.Department_name, &employee.Position_uuid, &employee.Position_name, &employee.Superior_uuid, &employee.Superior_name)
		employees = append(employees, employee)
	}

	return employees
}

// DeleteEmployee deletes an employee from the database using the DELETE SQL command.
// It takes an employee UUID string, http.ResponseWriter, and *http.Request as arguments.
func DeleteEmployee(employeeUuid string, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("DELETE FROM employee WHERE uuid = ?")
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
