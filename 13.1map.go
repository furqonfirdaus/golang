package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "furqon"
	//person["address"] = "cipulir"
	person := map[string]string{
		"name":    "furqon",
		"address": "Cipulir",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}

// hampirsama dengan array tetapi map ini bebas asalkan kata kuncinya berbeda
