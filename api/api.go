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

func readBody(r *http.Request) (b []byte) {
	b, err := ioutil.ReadAll(r.Body)
	hp.HandleErr(err)

	return
}

func login(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var lr LoginRequest
	err := json.Unmarshal(body, &r)
	hp.HandleErr(err)

	isValid := hp.ValidateReq([]md.Validation{
		{Value: lr.Username, Valid: "username"},
		{Value: lr.Password, Valid: "password"},
	})

	if !isValid {
		json.NewEncoder(w).Encode(md.ErrResponse{
			Message: "Wrong details",
		})
	}

	res := users.Login(&users.LoginParams{
		Username: lr.Username,
		Password: lr.Password,
	})

	if res["message"] != "login successful" {
		json.NewEncoder(w).Encode(md.ErrResponse{
			Message: "Wrong password or username",
		})
	}
	json.NewEncoder(w).Encode(res)
}

func register(w http.ResponseWriter, r *http.Request) {

	body := readBody(r)

	var rr RegisterRequest
	err := json.Unmarshal(body, &rr)
	hp.HandleErr(err)

	isValid := hp.ValidateReq([]md.Validation{
		{Value: rr.Username, Valid: "username"},
		{Value: rr.Email, Valid: "email"},
		{Value: rr.Password, Valid: "password"},
	})

	if !isValid {
		json.NewEncoder(w).Encode(md.ErrResponse{
			Message: "Wrong details",
		})
	}

	res := users.Register(&users.RegisterParams{
		Username: rr.Username,
		Email:    rr.Email,
		Password: rr.Password,
	})

	if res["message"] != "registration successful" {
		json.NewEncoder(w).Encode(md.ErrResponse{
			Message: "Wrong registration details",
		})
	}

	json.NewEncoder(w).Encode(res)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)

	json.NewEncoder(w).Encode(user)
}

func StartServer() {
	r := mux.NewRouter()
	r.Use(hp.PanicHandler)
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")

	fmt.Println("Server running....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
