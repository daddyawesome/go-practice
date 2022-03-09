package main

import (
	"fmt"
	"net/http"
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

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello World!!!")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/hello", helloWorld)
	fmt.Println("Starting the server om : 3000 ...")
	http.ListenAndServe(":3000", nil)
}
