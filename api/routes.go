package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartApi() {

	const webPort = ":8888"

	router := mux.NewRouter()

	//Register routes
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")

	fmt.Printf("Starting server on port %s", webPort)
	log.Fatal(http.ListenAndServe(":8888", router))

}
