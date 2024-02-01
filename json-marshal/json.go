package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	// convert to struct from json
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



	// convert to json from struct
	var mySlice []Person

	var m1 Person
	m1.FirstName ="Abhi"
	m1.LastName="shukla"
	m1.HairColor="black"
	m1.HasDog=false
	mySlice =append(mySlice, m1)

	var m2 Person
	m2.FirstName ="swapnil"
	m2.LastName="shukla"
	m2.HairColor="black"
	m2.HasDog=true
	mySlice =append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil{
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))

}
