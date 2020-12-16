package main

import (
	"net/http"
	"fmt"
	"restful/routes"
)

func main() {
	r := routes.NewRouter()
	fmt.Printf("Server Listening at : http://localhost:8080/\n")
	http.ListenAndServe(":8080", r)
	
}
