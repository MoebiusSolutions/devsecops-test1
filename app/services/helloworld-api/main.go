package main

import (
	"net/http"

	myhndlr "devsecops-test1/app/services/helloworld-api/handlers"
)

func main() {
	http.HandleFunc("/user", myhndlr.GreetUser)
	http.HandleFunc("/", myhndlr.Greet)
	http.ListenAndServe(":8080", nil)
}
