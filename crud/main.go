package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Id   int
	Name string
	City string
}

type Children struct {
	Id   int
	Name string
	Age  int
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := dbConn()
	//   // create database
	// _, err := db.Exec("CREATE DATABASE golang")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("created db")

	//   //create table
	// _, err := db.Exec("CREATE TABLE Children(Id int, Name varchar(20), Age varchar(20))")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("created table")

	//   //insert into table
	_, err := db.Exec("INSERT INTO Children VALUES(1, 'betuu', '08')")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(" data inserted")

	//   //select from table

	//    //Insert Into Database
	// sql := "INSERT INTO student(Name,City) VALUES ('swati','indore')"
	// res, err := db.Exec(sql)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// lastId, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("The last inserted row id: %d\n", lastId)

	//Select From Database
	// res, err := db.Query("SELECT * FROM student")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for res.Next() {
	// 	var student Student
	// 	err := res.Scan(&student.Id, &student.Name, &student.City)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%v\n", student)
	// }

	//Update From Database
	// sql := "UPDATE student SET Name = 'smita' WHERE Id = 1"
	// _, err := db.Exec(sql)
	// if err != nil {
	// 	panic(err.Error())
	// }

	//Delete From Database
	// sql := "DELETE FROM student WHERE Id = 3"
	// _, err := db.Exec(sql)
	// if err != nil {
	// 	panic(err.Error())
	// }
}
