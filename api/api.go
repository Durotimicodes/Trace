package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
	"github.com/durotimicodes/trace-backend/users"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	//Handle Login
	var formattedBody models.Login

	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Prepare response
	if login["message"] == "Login Successfull" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		//Handle error
		resp := models.ErrResponse{Message: "Invalid credentials"}
		json.NewEncoder(w).Encode(resp)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	//Handle Register
	var formattedBody models.Register

	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	//Prepare response
	if register["message"] == "All is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		//Handle error
		resp := models.ErrResponse{"Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}
