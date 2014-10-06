// reading https://golang.org/doc/effective_go.html

package main

import (
    "fmt"
)

const (
    Enone = 3
    Eio = 1
    Einval = 4
)

func main() {
    a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
    m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

    for i, v := range a {
        fmt.Printf("%d : %v\n", i, v)
    }
    fmt.Println("----------")

    for i, v := range s {
        fmt.Printf("%d : %v\n", i, v)
    }
    fmt.Println("----------")

    for k, v := range m {
        fmt.Printf("%v : %v\n", k, v)
    }
}
