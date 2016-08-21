package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// NewServer configures the web server
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler()).Methods("GET")

	n.UseHandler(router)
	return n
}

func helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		name := os.Getenv("NAME_TO_GREET")
		if len(name) == 0 {
			name = "stranger"
		}
		greeting := fmt.Sprintf("Hello, %s!", name)
		w.WriteHeader(200)
		w.Write([]byte(greeting))
	}
}
