package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "ProjetFinalSQL.db?_foreign_keys=on")
	fmt.Println("open/create DB:")
	checkErr(err)
	// defer close
	defer db.Close()

	dropTables := `
PRAGMA foreign_keys = OFF;
DROP TABLE IF EXISTS employee;
DROP TABLE IF EXISTS department;
DROP TABLE IF EXISTS team;
DROP TABLE IF EXISTS position;
DROP TABLE IF EXISTS leave;
DROP TABLE IF EXISTS employee_team;
	`
	_, err = db.Exec(dropTables)
	fmt.Println("drop tables:")
	checkErr(err)
	dropTables += ""

	createTables := `
CREATE TABLE employee(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY,  
	last_name VARCHAR(255), 
	first_name VARCHAR(255), 
	email VARCHAR(64), 
	phone_number VARCHAR(14), 
	department_uuid VARCHAR(255),
	position_uuid VARCHAR(255),
	superior_uuid VARCHAR(255),
	FOREIGN KEY (department_uuid) REFERENCES department(uuid),
	FOREIGN KEY (position_uuid) REFERENCES position(uuid),
	FOREIGN KEY (superior_uuid) REFERENCES employee(uuid));

CREATE TABLE department( 
	uuid VARCHAR(255) NOT NULL PRIMARY KEY,
	department_leader_uuid VARCHAR(255),  
	name VARCHAR(255),
	FOREIGN KEY (department_leader_uuid) REFERENCES employee(uuid)
	);

CREATE TABLE team(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY,
	team_leader_uuid VARCHAR(255),
	name VARCHAR(32) UNIQUE,
	FOREIGN KEY (team_leader_uuid) REFERENCES employee(uuid));

CREATE TABLE position(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY, 
	title VARCHAR(32), 
	salary INTEGER);

CREATE TABLE leave(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY,
	employee_uuid VARCHAR(255), 
	start_date DATETIME,
	end_date DATETIME,
	leave_type VARCHAR(255), 
	FOREIGN KEY (employee_uuid) REFERENCES employee(uuid));

CREATE TABLE employee_team(
	employee_uuid VARCHAR(255), 
	team_uuid VARCHAR(255), 
	PRIMARY KEY(employee_uuid, team_uuid), 
	FOREIGN KEY (employee_uuid) REFERENCES employee(uuid), 
	FOREIGN KEY (team_uuid) REFERENCES team(uuid));
	`

	_, err = db.Exec(createTables)
	fmt.Println("create tables:")
	checkErr(err)

	createTables += ""
	fmt.Println("Successfuly created the database!")

	inserts := `INSERT INTO employee (uuid, last_name, first_name, email, phone_number, department_uuid, position_uuid, superior_uuid) VALUES 
	('1', 'Doe', 'John', 'email', '06', '1', '1', '2'),
	('2', 'Smith', 'Jane', 'email2', '07', '1', '2', '1');
	
	INSERT INTO team (uuid, team_leader_uuid, name) VALUES 
	('1', '1', 'Accounting');
	
	INSERT INTO employee_team (employee_uuid, team_uuid) VALUES 
	('1', '1');
	
	INSERT INTO department (uuid, department_leader_uuid, name) VALUES 
	('1', '1', 'HR');
	
	INSERT INTO position (uuid, title, salary) VALUES 
	('1', 'Team Leader', 20000),
	('2', 'Department Leader', 35000);`

	_, err = db.Exec(inserts)
	fmt.Println("insert values:")
	checkErr(err)

	inserts += ""
	fmt.Println("Successfuly inserted values!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
