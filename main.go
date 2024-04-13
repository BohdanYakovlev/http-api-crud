package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"httpApiTest/api"
	_ "httpApiTest/api"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

const dbOpenConfig = "host=localhost port=5432 user=postgres password=152152 dbname=apitestusers sslmode=disable"
const selectAllUsers = `SELECT * FROM users`
const selectUserById = `SELECT username,phone FROM users WHERE id = $1`
const insertUser = `INSERT INTO users(username, phone) VALUES ($1,$2)`
const deleteUSerById = `DELETE FROM users WHERE id = $1`
const updateUserById = `UPDATE users SET username = $2, phone = $3 WHERE id = $1`

var db *sql.DB

func getUsers(w http.ResponseWriter, r *http.Request) {

	users := []api.User{}

	rows, err := db.Query(selectAllUsers)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var tempUser api.User
		err = rows.Scan(&tempUser.Id, &tempUser.Username, &tempUser.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, tempUser)
	}
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	row := db.QueryRow(selectUserById, id)

	var tempUser api.User
	tempUser.Id = &id
	err := row.Scan(&tempUser.Username, &tempUser.Phone)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(tempUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {

	username, _ := mux.Vars(r)["username"]
	phone, _ := mux.Vars(r)["phone"]
	_, err := db.Exec(insertUser, username, phone)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := db.Exec(deleteUSerById, id)

	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	username, _ := mux.Vars(r)["username"]
	phone, _ := mux.Vars(r)["phone"]

	_, err := db.Exec(updateUserById, id, username, phone)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	var err error
	db, err = sql.Open("postgres", dbOpenConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{username}/{phone}", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}/{username}/{phone}", updateUser).Methods("PUT")

	reg := httptest.NewRequest("GET", "/users", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("POST", "/users/testname7/testphone7", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("DELETE", "/users/6", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("PUT", "/users/7/testname7/testphone7", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

	reg = httptest.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, reg)
	fmt.Println(recorder.Body.String())

}
