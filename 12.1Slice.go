package main

import "fmt"

func main() {
	names := [...]string{"januari", "februari", "maret", "april", "mei", "juni"}
	slice := names[4:6]
	fmt.Println(slice)

	slice1 := names[:4]
	fmt.Println(slice1)

	slice2 := names[2:]
	fmt.Println(slice2)

	//copy slice
	fromSlice := names[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
