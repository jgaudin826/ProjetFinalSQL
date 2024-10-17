package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on") // file:./threadcore.db?_foreign_keys=on
	fmt.Println("open/create DB:")
	checkErr(err)
	// defer close
	defer db.Close()

	dropTables := `
PRAGMA foreign_keys = OFF;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS like;
DROP TABLE IF EXISTS community;
DROP TABLE IF EXISTS user_community;
DROP TABLE IF EXISTS friend;
DROP TABLE IF EXISTS user;
	`

	// not used for now
	//DROP TABLE IF EXISTS message;
	//DROP TABLE IF EXISTS groupchat;
	//DROP TABLE IF EXISTS user_groupchat;
	//DROP TABLE IF EXISTS friend_request;

	_, err = db.Exec(dropTables)
	fmt.Println("drop tables:")
	checkErr(err)
	dropTables += ""

	createTables := `
CREATE TABLE user(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
	uuid VARCHAR(255),  
	profile VARCHAR(255), 
	banner VARCHAR(255), 
	email VARCHAR(64), 
	username VARCHAR(20), 
	password VARCHAR(255));

CREATE TABLE community(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
	profile VARCHAR(255), 
	banner VARCHAR(255), 
	name VARCHAR(32), 
	description VARCHAR(255), 
	user_id INTEGER, 
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE);

CREATE TABLE friend(
	user_id INTEGER,
	friend_id INTEGER, 
	PRIMARY KEY(user_id, friend_id), 
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, 
	FOREIGN KEY (friend_id) REFERENCES user(id) ON DELETE CASCADE);

CREATE TABLE post(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
	uuid VARCHAR(255),
	title VARCHAR(32), 
	content VARCHAR(500), 
	media VARCHAR(255), 
	media_type VARCHAR(255),
	user_id INTEGER, 
	community_id INTEGER, 
	created TIMESTAMP, 
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, 
	FOREIGN KEY (community_id) REFERENCES community(id) ON DELETE CASCADE);

CREATE TABLE comment(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
	user_id INTEGER, 
	post_id INTEGER, 
	comment_id INTEGER, 
	content VARCHAR(255), 
	created TIMESTAMP, 
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, 
	FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE, 
	FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE);

CREATE TABLE like(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
	rating INTEGER, 
	comment_id INTEGER, 
	post_id INTEGER, 
	user_id INTEGER, 
	FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE, 
	FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE, 
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE);

CREATE TABLE user_community(
	user_id INTEGER, 
	community_id INTEGER, 
	PRIMARY KEY(user_id, community_id), 
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, 
	FOREIGN KEY (community_id) REFERENCES community(id) ON DELETE CASCADE);
	`

	// not used for now :
	//CREATE TABLE message(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user_id INTEGER, groupchat_id INTEGER, friend_id INTEGER, message VARCHAR(255), sent TIMESTAMP, FOREIGN KEY (user_id) REFERENCES friend(user_id) ON DELETE CASCADE, FOREIGN KEY (groupchat_id) REFERENCES groupchat(id) ON DELETE CASCADE, FOREIGN KEY (friend_id) REFERENCES friend(friend_id) ON DELETE CASCADE);
	//CREATE TABLE user_groupchat(user_id INTEGER, groupchat_id INTEGER, PRIMARY KEY(user_id, groupchat_id), FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, FOREIGN KEY (groupchat_id) REFERENCES groupchat(id) ON DELETE CASCADE);
	//CREATE TABLE friend_request(user_id INTEGER, request_id INTEGER, PRIMARY KEY(user_id, request_id), FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, FOREIGN KEY (request_id) REFERENCES user(id) ON DELETE CASCADE);
	//CREATE TABLE groupchat(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name VARCHAR(32));

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
