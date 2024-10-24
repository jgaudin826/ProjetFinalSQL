package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Department struct {
	Uuid                   string
	Department_leader_uuid string
	Name                   string
}

type DepartmentInfo struct {
	Uuid                   string
	Department_leader_uuid string
	Department_leader_name string
	Name                   string
}

// AddDepartment opens the database connection and adds a department to it using the INSERT INTO SQL command.
// It takes a Department struct, http.ResponseWriter, and *http.Request as arguments.
func AddDepartment(department Department, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO department (uuid, department_leader_uuid, name) VALUES (?, ?, ?)")
	query.Exec(department.Uuid, department.Department_leader_uuid, department.Name)
	CheckErr(err2, w, r)
	defer query.Close()
}

// GetDepartmentByUuid retrieves a department by its UUID using the SELECT * FROM SQL command.
// It takes a UUID string, http.ResponseWriter, and *http.Request as arguments, and returns a Department struct.
func GetDepartmentByUuid(uuid string, w http.ResponseWriter, r *http.Request) Department {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM department WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	department := Department{}

	for rows.Next() {
		rows.Scan(&department.Uuid, &department.Department_leader_uuid, &department.Name)
	}

	return department
}

/*
!GetdepartmentByName function is used to get a department by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return a department type.
*/
func GetDepartmentByName(departmentName string, w http.ResponseWriter, r *http.Request) Department {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM department WHERE name = '" + departmentName + "'")
	defer rows.Close()

	department := Department{}

	for rows.Next() {
		rows.Scan(&department.Uuid, &department.Department_leader_uuid, &department.Name)
	}

	return department
}

/*
!GetEmployeesByTeam function open data base and get users on a community by using the SELECT * FROM sql command she take as argument an int type and a writer and request and return a slice of User.
*/
func GetEmployeesByDepartment(departmentUuid string, w http.ResponseWriter, r *http.Request) []EmployeeInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT e.uuid, e.last_name, e.first_name, e.email, e.phone_number, e.department_uuid, d.name as department_name, e.position_uuid, p.title as position_name, e.superior_uuid, s.first_name || ' ' || s.last_name as superior_name FROM employee e LEFT JOIN department d ON e.department_uuid = d.uuid LEFT JOIN position p ON e.position_uuid = p.uuid LEFT JOIN employee s ON e.superior_uuid = s.uuid WHERE e.department_uuid='" + departmentUuid + "'")
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

func GetAllDepartments(w http.ResponseWriter, r *http.Request) []DepartmentInfo {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	rows, err := db.Query("SELECT d.uuid, d.department_leader_uuid, e.first_name || ' ' || e.last_name as department_leader_name, d.name FROM department d LEFT JOIN employee e ON d.department_leader_uuid = e.uuid")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	departmentList := make([]DepartmentInfo, 0)

	for rows.Next() {
		department := DepartmentInfo{}
		err = rows.Scan(&department.Uuid, &department.Department_leader_uuid, &department.Department_leader_name, &department.Name)
		CheckErr(err, w, r)

		departmentList = append(departmentList, department)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return departmentList
}

// UpdateDepartmentInfo updates a department's information using the UPDATE SQL command.
// It takes a Department struct, http.ResponseWriter, and *http.Request as arguments.
func UpdateDepartmentInfo(department Department, w http.ResponseWriter, r *http.Request) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the database at the end of the function
	defer db.Close()

	query, err := db.Prepare("UPDATE department SET department_leader_uuid = ?, name = ? WHERE uuid = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(department.Department_leader_uuid, department.Name, department.Uuid)
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
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
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
