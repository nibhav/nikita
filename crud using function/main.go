package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	Id      int
	Name    string
	Address string
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cartoon")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// func CreateDB(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	_, err := db.Exec("CREATE DATABASE cartoon")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Database created")
// 	fmt.Fprintln(w, "db created")
// }

func CreateTable(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	_, err := db.Exec("CREATE TABLE tree(Id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, Name varchar(20), Age varchar(20))")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("created table TREE")
	fmt.Fprintln(w, "TREE table created successfully ", r)
}

// func Insert(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	_, err := db.Exec("INSERT INTO mickey(Name, Age) values('tom','100')")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Data Inserted Successfully ")
// }

// func Show(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	rows, err := db.Query("SELECT * FROM student")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		var city string
// 		rows.Scan(&id, &name, &city)
// 		//fmt.Printf("Id: %d Name: %s City: %s", id, name, city)
// 		fmt.Fprintf(w, "Id %d:", id)
// 		fmt.Fprintf(w, "Name %s:", name)
// 		fmt.Fprintln(w, "City :", city)
// 	}
// }

// func Update(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	update, err := db.Prepare("UPDATE student SET name = swati WHERE id = 3")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_, err = update.Exec()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Data Updated Successfully ", update)
// }

func main() {
	//log.Println("Server started on: http://localhost:9000")
	//http.HandleFunc("/create", Create)
	//http.HandleFunc("/insert", Insert)
	//http.HandleFunc("/show", Show)
	//http.HandleFunc("/createdb", CreateDB)
	http.HandleFunc("/createtable", CreateTable)
	fmt.Println(http.ListenAndServe(":4020", nil))

}
