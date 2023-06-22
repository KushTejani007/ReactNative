// package main

// import (
// 	"net/http"
// 	"fmt"
// 	"github.com/go-sql-driver/mysql"
// 	"gorm.io/gorm"
// )

// var err error

// const DNS = ("root:root@tcp(localhost:3306)/transactions")

// type Users struct {
// 	gorm.Model
// 	SenderName   string `json:"t1"`
// 	ReceiverName string `json:"t2"`
// 	Amount       int    `json:"t3"`
// 	Date         int    `json:"t4"`
// }

// func InitalizeMigration() {

// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {

// }

// func CreateUsers(w http.ResponseWriter, r *http.Request) {

// }

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace <username>, <password>, <hostname>, and <database> with your MySQL connection details
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/transactions")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database!")

	insquery := "INSERT INTO transaction(lenders_name,senders_name, amount, date) VALUES ( ?, ?, ?, ?)"

	lname := "lender4"
	rname := "receiver4"
	amt := 50000
	date := "31-01-12"
	rows, err := db.Query(insquery, lname, rname, amt, date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
