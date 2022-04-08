package main

import (
	"fmt"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Welcome To AWESOMENESS!!!!</h1>")
	case "/daddy":
		fmt.Fprint(w, "<h1>I am daddy AWESOME</h1>")
	default:
		fmt.Fprint(w, "Error Daw!!!")
	}

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>to get in touch hmmm<p>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attemp")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not Time OUt")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/timeout", timeout)
	fmt.Println("Starting the server on : 3000 ...")
	http.ListenAndServe(":3000", nil)
}
