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
	router.HandleFunc("/user/{id}", getUserHandler).Methods("GET")

	fmt.Printf("Starting server on port %s", webPort)
	log.Fatal(http.ListenAndServe(":8888", router))

}
