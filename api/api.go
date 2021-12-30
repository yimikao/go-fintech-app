package api

import (
	"encoding/json"
	"fmt"
	hp "go-fintech-app/helpers"
	md "go-fintech-app/models"
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

type RegisterRequest struct {
	Username string
	Email    string
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

	isValid := hp.ValidateReq([]md.Validation{
		{Value: lr.Username, Valid: "username"},
		{Value: lr.Password, Valid: "password"},
	})

	if !isValid {
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "Wrong details",
		})
	}

	res := users.Login(&users.LoginParams{
		Username: lr.Username,
		Password: lr.Password,
	})

	if res["message"] != "login successful" {
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "Wrong password or username",
		})
	}
	json.NewEncoder(w).Encode(res)
}

func register(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	hp.HandleErr(err)

	var rr RegisterRequest
	err = json.Unmarshal(body, &rr)

	isValid := hp.ValidateReq([]md.Validation{
		{Value: rr.Username, Valid: "username"},
		{Value: rr.Email, Valid: "email"},
		{Value: rr.Password, Valid: "password"},
	})

	if !isValid {
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "Wrong details",
		})
	}

	res := users.Register(&users.RegisterParams{
		Username: rr.Username,
		Email:    rr.Email,
		Password: rr.Password,
	})

	if res["message"] != "registration successful" {
		json.NewEncoder(w).Encode(ErrResponse{
			Message: "Wrong registration details",
		})
	}

	json.NewEncoder(w).Encode(res)
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")

	fmt.Println("Server running....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
