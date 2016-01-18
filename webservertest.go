package main

import (
	"net/http"
	"encoding/json"
	"time"
	"log"
)

type Param struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Numbers []int `json:"numbers"`
}

func handler(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("This is an example server. \n"))
}

func postParamHandler(w http.ResponseWriter, r *http.Request)	{
	params := Param{
		Name: "nabil",
		Email: "nabil.sooz@gmail.com",
		Numbers: []int{111,222,333},
	}

	p, err := json.Marshal(params)

	if err != nil	{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(p)
}

func getTimeHandler(w http.ResponseWriter, r *http.Request)	{
	t, err := json.Marshal(time.Now())

	if err != nil	{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(t)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/params", postParamHandler)
	http.HandleFunc("/time", getTimeHandler)

	err := http.ListenAndServe(":3030", nil)
	if err != nil	{
		log.Fatal("ListenAndServe: ", err)
	}
}
