package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Payload ...
type Payload struct {
	Stuff Data
}

// Data ...
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

// Fruits ...
type Fruits map[string]int

// Vegetables ...
type Vegetables map[string]int

// serveRest ...
func serveRest(w http.ResponseWriter, req *http.Request) {
	response, error := getJSONResponse()
	if error != nil {
		panic(error)
	}
	fmt.Fprintf(w, string(response))
}

// getJsonResponse ...
func getJSONResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:12345", nil)
}
