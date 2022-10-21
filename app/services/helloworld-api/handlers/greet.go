package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type WrapperStruct struct {
	Name string
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now().UTC().Format("01-02-2006"))
}

func GreetUser(w http.ResponseWriter, r *http.Request) {
	pu := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello World! %s", pu)
}
