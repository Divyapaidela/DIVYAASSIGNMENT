package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Divya@135@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"

type Emp struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	Dpt  string `json:"dpt"`
}

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", GetEmps).Methods("GET")
	r.HandleFunc("/employees/{id}", GetEmp).Methods("GET")
	r.HandleFunc("/employees", CreateEmp).Methods("POST")
	r.HandleFunc("/employees/{id}", UpdateEmp).Methods("PUT")
	r.HandleFunc("/employees/{id}", DeleteEmp).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}

func InitialMigration() {
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Emp{})

}

func GetEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp []Emp
	DB.Find(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var emp Emp
	DB.First(&emp, params["id"])
	json.NewEncoder(w).Encode(emp)
}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Emp
	json.NewDecoder(r.Body).Decode(&emp)
	DB.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var emp Emp
	DB.First(&emp, params["id"])
	json.NewDecoder(r.Body).Decode(&emp)
	DB.Save(&emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var emp Emp
	DB.Delete(&emp, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
