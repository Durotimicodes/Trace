package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/gorilla/mux"
)

func StartApi() {

	const webPort = ":8888"
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)

	//Registered routes
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/transaction", transactionHandler).Methods("POST")
	router.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/transactions/{userID}", getMyTransaction).Methods("GET")
	fmt.Printf("Trace Bank App working on port %s", webPort)
	log.Fatal(http.ListenAndServe(":8888", router))

}
