package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	//http get Request
	url := "http://localhost:12345"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var p Payload

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Stuff.Fruit, "\n", p.Stuff.Veggies)
}
