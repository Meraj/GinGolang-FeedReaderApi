package app

import (
	"database/sql"
 _ "github.com/go-sql-driver/mysql"

	)


func Database() *sql.DB {
	connection := "root:@tcp(127.0.0.1:3306)/test"
	//	db, err := sql.Open("mysql", dbInfo{}.dbUser+":"+dbInfo{}.dbPass+"@tcp("+dbInfo{}.dbHost+")/"+dbInfo{}.dbName)
		db, err := sql.Open("mysql", connection)
	Migration(db)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Query(query_str string) (*sql.Rows, error) {
	db := Database()
	return db.Query(query_str)
}


func Migration(db *sql.DB){
	var queries []string
	queries = append(queries,"CREATE TABLE IF NOT EXISTS users (id int (255) NOT NULL AUTO_INCREMENT,username VARCHAR (255) NOT NULL,email VARCHAR (255) NOT NULL,password MEDIUMTEXT NOT NULL,account_status int (5) NOT NULL,token VARCHAR (255),api_token VARCHAR (255),  PRIMARY KEY (id));")
	queries = append(queries,"CREATE TABLE IF NOT EXISTS feeds (id int (255) NOT NULL AUTO_INCREMENT,name VARCHAR (255) NOT NULL,url VARCHAR (255) NOT NULL,content_selector VARCHAR (255),  PRIMARY KEY (id));")
	queries = append(queries,"CREATE TABLE IF NOT EXISTS contents (id int (255) NOT NULL AUTO_INCREMENT,url VARCHAR (255),title VARCHAR (255),body LONGTEXT,source_id int (255),  PRIMARY KEY (id));")
	for i := range queries {
		db.Exec(queries[i])
	}
}