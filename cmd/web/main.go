package main

import (
	"fmt"
	"github.com/codykoscielski/learn-go/pkg/handlers"
	"net/http"
)

const portNumber = ":6969"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Start app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
