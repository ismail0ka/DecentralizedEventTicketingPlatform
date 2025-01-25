package main

import (
	"fmt"
	"log"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"database/sql"
)

var db *gorm.DB;

func main() {
	db = connectToDB("devent");
}

func connectToDB(db_name string) * gorm.DB {
	conn_string := $env.CONN_STRING + db_name;
	sqldb, err := sql.Open("mysql", conn_string);
	gormdb, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqldb}),
		&gorm.Config{});
	if err != nil {
		log.Fatal("Failed to connect to database:", err);
	}
	return gormdb;
}
