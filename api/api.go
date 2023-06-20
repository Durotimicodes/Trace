package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
	"github.com/durotimicodes/trace-backend/useraccounts"
	"github.com/durotimicodes/trace-backend/users"
	"github.com/gorilla/mux"
)

func readBody(r *http.Request) []byte {

	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {

	//Prepare response
	if call["message"] == "All is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		//Handle error
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Read body
	body := readBody(r)

	//Handle Login
	var formattedBody models.Login

	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Prepare response
	apiResponse(login, w)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//Read body
	body := readBody(r)

	//Handle Register
	var formattedBody models.Register

	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	//Prepare response
	apiResponse(register, w)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	auth := r.Header.Get("Authorization")

	//Handle Register
	var formattedBody models.TransactionBody

	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	transaction := useraccounts.Transaction(formattedBody.UserId, formattedBody.From, formattedBody.To, formattedBody.Amount, auth)

	//Prepare response
	apiResponse(transaction, w)

}
