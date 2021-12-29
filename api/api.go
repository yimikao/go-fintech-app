package api

import (
	"encoding/json"
	"fmt"
	hp "go-fintech-app/helpers"
	"go-fintech-app/users"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginRequest struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	hp.HandleErr(err)

	var lr LoginRequest
	err = json.Unmarshal(body, &r)
	hp.HandleErr(err)

	res := users.Login(lr.Username, lr.Password)

	if res["message"] != "login successful" {
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "Wrong password or username",
		})
	}
	json.NewEncoder(w).Encode(res)
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")

	fmt.Println("Server running....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
