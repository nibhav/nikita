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
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// func Create(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	emp, err := db.Prepare("CREATE TABLE emplyee(ID int,Name string,Address string)")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	_, err = emp.Exec(emp)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Employee table created", 302)
// }

// func Insert(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	insert, err := db.Prepare("INSERT INTO student(name, city) values('vaibhav','bhusawal')")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_, err = insert.Exec()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Data Inserted Successfully ", insert)
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
	http.HandleFunc("/create", Create)
	//http.HandleFunc("/index", Insert)
	//http.HandleFunc("/show", Show)
	fmt.Println(http.ListenAndServe(":9000", nil))

}
