package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "furqon"
	newSlice[1] = "furqon"
	//newSlice[2] = "furqon" (ini ga bisa harus pake append, soalnya nilainya udh ditentuin di atas yang 2 itu)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//contoh pake append
	newSlice2 := append(newSlice, "furqon")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	//coba array
	newSlice2[0] = "firdaus"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//copy slice
	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
