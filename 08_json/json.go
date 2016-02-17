package main

import "fmt"
import "encoding/json"

type Person struct {
	Name string
	Age  int
}

func main() {
	//4 main functions that are associated with json files
	//Encoder: from go to json
	//Decoder: from json to go
	//-------------uses reader/writer
	//Marshal: from go to json
	//Unmarshal: from json to go
	//-------------uses []bytes of data
	p := Person{"Pheng", 10}
	bs, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	//marshall example

	fmt.Println(string(bs))

	var q Person
	json.Unmarshal(bs, &q)

	//unmarshall example

	fmt.Println(q)
}
