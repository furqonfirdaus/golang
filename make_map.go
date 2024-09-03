package main

import "fmt"

func main() {
	book := make(map[string]string)
	book["title"] = "buku go-lang"
	book["author"] = "Furqon Firdaus"
	book["wrong"] = "ups"

	delete(book, "wrong") //ini akan menghapus kata wrong: ups

	fmt.Println(book)
}
