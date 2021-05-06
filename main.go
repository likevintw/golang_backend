package main

import (
	"fmt"
	"http_template/server"
	"net/http"
)

func main() {
	// RESTful API
	fmt.Println(("Set Up RESTful Server"))

	//
	http.Handle("/user", new(server.HttpServerJson))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
