package main

import (
	"net/http"

	route "github.com/wei840222/todo-api/route"
)

func main() {
	r := route.NewRouter()

	http.ListenAndServe(":8080", r)
}
