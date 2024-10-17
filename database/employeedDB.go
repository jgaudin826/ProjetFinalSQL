package database

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Uuid     string
	Profile  string
	Banner   string
	Email    string
	Username string
	Password string
}

/*
!AddUser function open data base and add an user to it with the INSERT INTO sql command she take as argument an User type and a writer and request.
*/
func AddUser(user User, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err2 := db.Prepare("INSERT INTO user (uuid, profile, banner, email, username, password) VALUES (?, ?, ?, ?, ?, ?)")
	query.Exec(user.Uuid, user.Profile, user.Banner, user.Email, user.Username, user.Password)
	CheckErr(err2, w, r)
	defer query.Close()

	query2, err3 := db.Prepare("INSERT INTO friend (user_id, friend_id) VALUES (?, ?)")
	query2.Exec(user.Id, user.Id)
	CheckErr(err3, w, r)
	defer query2.Close()
}

/*
!GetUserByUuid function is used to get a user by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return an User type.
*/
func GetUserByUuid(uuid string, w http.ResponseWriter, r *http.Request) User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM user WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

/*
!UpdateUserInfo function is used to update user inforamation by using UPDATE sql command. She take as argument a user type, a writer, a request.
*/
func UpdateUserInfo(user User, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE user SET uuid = ?, profile = ?, banner = ?, username = ?, email = ?, password = ? WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(user.Uuid, user.Profile, user.Banner, user.Username, user.Email, user.Password, user.Id)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

/*
!DeleteUser function is used to delete user by using DELETE sql command. She take as argument an int, a writer, a request.
*/
func DeleteUser(userId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM user WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(userId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was deleted")
	}
}
