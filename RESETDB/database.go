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
	department_uuid INTEGER,
	position_uuid INTEGER,
	superior_uuid INTEGER,
	FOREIGN KEY (department_uuid) REFERENCES department(uuid)ON DELETE CASCADE,
	FOREIGN KEY (position_uuid) REFERENCES position(uuid)ON DELETE CASCADE,
	FOREIGN KEY (superior_uuid) REFERENCES employee(uuid)ON DELETE CASCADE);

CREATE TABLE department( 
	uuid VARCHAR(255) NOT NULL PRIMARY KEY, 
	name VARCHAR(255),
	department_leader_uuid INTEGER, 
	FOREIGN KEY (department_leader_uuid) REFERENCES employee(uuid) ON DELETE CASCADE
	);

CREATE TABLE team(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY,
	team_leader_uuid INTEGER, 
	FOREIGN KEY (team_leader_uuid) REFERENCES employee(uuid) ON DELETE CASCADE);

CREATE TABLE position(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY, 
	title VARCHAR(32), 
	salary INTEGER);

CREATE TABLE leave(
	uuid VARCHAR(255) NOT NULL PRIMARY KEY,
	employee_uuid INTEGER, 
	start_date DATETIME,
	end_date DATETIME,
	leave_type VARCHAR(255), 
	FOREIGN KEY (employee_uuid) REFERENCES employee(uuid) ON DELETE CASCADE);

CREATE TABLE employee_team(
	employee_uuid INTEGER, 
	team_uuid INTEGER, 
	PRIMARY KEY(employee_uuid, team_uuid), 
	FOREIGN KEY (employee_uuid) REFERENCES employee(uuid) ON DELETE CASCADE, 
	FOREIGN KEY (team_uuid) REFERENCES team(uuid) ON DELETE CASCADE);
	`

	_, err = db.Exec(createTables)
	fmt.Println("create tables:")
	checkErr(err)

	createTables += ""
	fmt.Println("Successfuly created the database!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
