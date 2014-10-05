// more than http://tour.golang.org/#60

package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

type defaultHandler struct{}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, fmt.Sprintf("%s %s %s", s.Greeting, s.Punct, s.Who))
}

func (h defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("undefined:", r.URL.Path)
    fmt.Fprint(w, "4_____0_____4")
}

func main() {
    http.Handle("/", defaultHandler{})
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
