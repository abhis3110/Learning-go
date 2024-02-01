package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `[
		{
			"first_name":"Clark",
			"last_name": "Kent",
			"hair_color":"black",
			"has_dog": true
		},
		{
			"first_name":"satyam",
			"last_name": "garg",
			"hair_color":"grey",
			"has_dog": false
		}
	]`

	var unmarshalled []Person // slice bcz json data support

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error unmarshalling json", err)
	}
	log.Printf("Unmarshalled %v", unmarshalled)

}
