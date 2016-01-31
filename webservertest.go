package main

import (
	"net/http"
	"encoding/json"
	"time"
	"log"
	"strconv"
)

type Param struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Numbers int `json:"numbers"`
}

func handler(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("This is an example server. \n"))
}

func postParamHandler(w http.ResponseWriter, r *http.Request)	{
	r.ParseForm()

	nums, err := strconv.Atoi(r.FormValue("numbers"))

	if err != nil	{
		http.Error(w, error.Error(err), 500)
		return
	}

	p := &Param{r.FormValue("name"), r.FormValue("email"), nums*2}

	n, err := json.Marshal(p)

	if err != nil	{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(n)
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

	err := http.ListenAndServe(":1112", nil)
	if err != nil	{
		log.Fatal("ListenAndServe: ", err)
	}
}
