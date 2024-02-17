package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":6969"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(4, 5)
	fmt.Fprintf(w, fmt.Sprintf("this is a thing, %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}

	result := x / y
	return result, nil
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Start app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
