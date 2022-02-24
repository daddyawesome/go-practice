package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome To AWESOMENESS!!!!</h1>")
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
