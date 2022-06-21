package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

	"gorm.io/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect() error {
	var err error
	Connector, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	Connector.AutoMigrate(&Person{})
	// Connector.Exec("CREATE TABLE IF NOT EXISTS personns(id INTEGER,FirstName TEXT, LastName TEXT, age INTEGER)")
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}
func main() {
	err := Connect()
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "Hello World!")
		} else if r.Method == "POST" {

			var u1 Person

			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println(body)
			json.Unmarshal(body, &u1)
			fmt.Print(u1)

			Connector.Create(&u1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, u1)

		}

	})
	http.ListenAndServe(":80", nil)
	log.Println("Starting the HTTP server on port 8090")

}

type Person struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
