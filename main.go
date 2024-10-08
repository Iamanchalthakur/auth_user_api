package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tetsing/controller"
)

func main() {
	connStr := "user=postgres password=123456 dbname=testing host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	http.HandleFunc("/create-user", controller.CreateUser(db))
	http.HandleFunc("/login", controller.Login(db))
	http.HandleFunc("/get-users", controller.GetAllUsers(db))
	http.HandleFunc("/delete-user", controller.DeleteAllUser(db))
	http.HandleFunc("/update-user/{id}", controller.UpdateUser(db))

	log.Println("Server started on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
