package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

//Create an api with jwt authentication
//and user management
func main() {
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

//create user model	struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//creat user model array
type Users []User

//creater a new router
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", Index)
	mux.HandleFunc("/api/users", CreateUser)
	mux.HandleFunc("/api/users/login", Login)
	return mux
}

//IOndex
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the Go JWT Tutorial"}`))
}

//function to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var users Users

	//get the body of our POST request
	//returns a byte slice
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//we unmarshal our byte slice into
	//a user struct
	json.Unmarshal(body, &user)

	//add our user to our user array
	users = append(users, user)

	//marshal/convert our users to json
	outgoingJSON, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	//write our users in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.Write(outgoingJSON)
}

//add jwt authentication middleware
//to our login route
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	//var hmacSampleSecret []byte
	//get the body of our POST request
	//returns a byte slice
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		panic(err)
	}

	//we unmarshal our byte slice into
	//a user struct
	json.Unmarshal(body, &user)
	//check to see if the user exists
	//in our in-memory map
	if user.Username == "admin" && user.Password == "password" {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"admin":    true,
			"username": user.Username,
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(token)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(tokenString))
		w.WriteHeader(200)
	} else {
		//if the user doesn't exist return a http status code
		w.WriteHeader(http.StatusUnauthorized)
	}
}
