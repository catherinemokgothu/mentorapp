package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"com.cathy.mentor-backend/helper"
	"com.cathy.mentor-backend/users"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username                  string
	Password                  string
	Email                     string
	Role                      string
	Firstname                 string
	Lastname                  string
	City                      string
	Country                   string
	Contactnumber             string
	Leveltechnicalexpirience  string
	Currenttechlevel          string
	Skillsoffersmentoringdesc string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	//Create body
	body, err := ioutil.ReadAll(r.Body)
	helper.HandleError(err)

	//Handle login
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helper.HandleError(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Prepare response
	if login["message"] == "Success" {
		response := login
		json.NewEncoder(w).Encode(response)
	} else {
		response := ErrResponse{Message: "Incorrect username or password"}
		json.NewEncoder(w).Encode(response)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	//Create body
	body, err := ioutil.ReadAll(r.Body)
	helper.HandleError(err)

	//Handle login
	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helper.HandleError(err)
	register := users.Register(formattedBody.Username, formattedBody.Password, formattedBody.Email, formattedBody.Role, formattedBody.Firstname, formattedBody.Lastname, formattedBody.City, formattedBody.Country, formattedBody.Contactnumber, formattedBody.Currenttechlevel, formattedBody.Leveltechnicalexpirience, formattedBody.Skillsoffersmentoringdesc)

	//Prepare response
	if register["message"] == "Success" {
		response := register
		json.NewEncoder(w).Encode(response)
	} else {
		response := ErrResponse{Message: "Failed to register user"}
		json.NewEncoder(w).Encode(response)
	}
}

func LoginApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println(("Application started on port :8080"))
	log.Fatal(http.ListenAndServe(":8080", router))

}
