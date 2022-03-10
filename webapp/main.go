package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Welcome To AWESOMENESS!!!!</h1>")
	case "/daddy":
		fmt.Fprint(w, "<h1>I am daddy AWESOME</h1>")
	default:
		fmt.Fprint(w, "Error Daw!!!")
	}

}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attemp")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not Time OUt")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/timeout", timeout)
	fmt.Println("Starting the server om : 3000 ...")
	http.ListenAndServe(":3000", nil)
}
