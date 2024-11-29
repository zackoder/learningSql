package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/controllers"
	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err_db := sql.Open("sqlite3", "./data.db")
	if err_db != nil {
		fmt.Println(err_db)
		return
	}
	defer db.Close()
	models.Create_table(db)
	
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/login",controllers.Login)
	fmt.Println("server runing at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err, "we cant serve")
		return
	}
}