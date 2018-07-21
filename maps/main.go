package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	printMap(colors)

	var diffColors map[string]string
	printMap(diffColors)

	makeColors := make(map[string]string)
	makeColors["red"] = "#ff0000"
	makeColors["white"] = "#ffffff"
	makeColors["black"] = "#000000"
	printMap(makeColors)

	delete(makeColors, "red")
	printMap(makeColors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Key: ", key, " Value: ", value)
	}
	fmt.Println("-----------")
}
