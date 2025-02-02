package main

import (
	"fmt"
	"log"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"database/sql"
	"github.com/gorilla/mux"
)

var db *gorm.DB;

func main() {
	router := mux.NewRouter();
	http.ListenAndServe(":8080", router);
}

func connectToDB(db_name string) * gorm.DB {
	conn_string := "mysql:toor:@tcp(localhost:3306)/" + db_name;
	sqldb, err := sql.Open("mysql", conn_string);
	gormdb, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqldb}),
		&gorm.Config{});
	if err != nil {
		log.Fatal("Failed to connect to database:", err);
	}
	return gormdb;
}
