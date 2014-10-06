package main

import (
    "fmt"
)

type User struct {
    //                    Tags
	ID     int            "json:id"
	Name   string         "json:username"
	Email  string         "json:email"
	First  string         "json:first"
	Last   string         "json:last"
}

func main() {
    user := User{}
    fmt.Println(user)
}
