package main

import "fmt"

func main() {
	mySlice := []string{"hello", "there", "slice", "vs", "structs"}

	updateSlice(mySlice)

	fmt.Println(mySlice) // [bye there slice vs structs]
}

func updateSlice(a []string) {
	a[0] = "bye"
}

// slice contains ptr to head, capacity, length
// arrays are of fixed length. Slice can grow and shrink
