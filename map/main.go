package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " -> " + hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	printMap(colors)
}
