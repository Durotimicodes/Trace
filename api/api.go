package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type ErrResponse struct {
	Message string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	//Handle Login
	var formattedBody Login

	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Prepare response
	if login["message"] == "Login Successfull" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		//Handle error
		resp := ErrResponse{Message: "Invalid credentials"}
		json.NewEncoder(w).Encode(resp)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	//Handle Register
	var formattedBody Register

	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	//Prepare response
	if register["message"] == "All is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		//Handle error
		resp := ErrResponse{"Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}



func StartApi() {

	const webPort = ":8888"

	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")

	fmt.Printf("Starting server on port %s", webPort)
	log.Fatal(http.ListenAndServe(":8888", router))

}
