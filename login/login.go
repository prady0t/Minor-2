package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func login(w http.ResponseWriter, e *http.Request){
//sample login check if Jenkins works take 4
	person := Person{
		Name:  "Pradyot",
		Age:   22,
		Email: "prady0t@example.com",
	}

	jsonData, err := json.Marshal(person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
}

func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", login)
	fmt.Println("Serving at port 3001")
	err := http.ListenAndServe(":3001", nil)
	if err != nil{
		fmt.Println("Error Occured!")
	}
	
}