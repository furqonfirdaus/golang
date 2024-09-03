package main

import "fmt"

func main() {
	type noKTP string

	var ktpFurqon noKTP = "2002430230"

	var contoh string = "12102001200"
	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpFurqon)
	fmt.Println(contohKtp)
}
