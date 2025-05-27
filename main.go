package main

import (
	"assignment-go-day-one/ping"
	"assignment-go-day-one/reverse"
	"net/http"
)

func main() {
	ping.Init()
	reverse.Init()
	http.ListenAndServe(":8080", nil)
}
