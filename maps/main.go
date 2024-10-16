package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"Royal Blue":   "#4169E1",
		"Crimson Red":  "#DC143C",
		"Forest Green": "#228B22",
		"Goldenrod":    "#DAA520",
		"Slate Gray":   "#708090",
	}

	colors["Black"] = "#000000"
	delete(colors, "Goldenrod")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s : %s\n", color, hex)
	}
}
