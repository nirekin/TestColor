package main

import (
	"log"

	"github.com/fatih/color"
)

func main() {

	baseColors := []color.Attribute{color.FgRed, color.FgGreen, color.FgYellow, color.FgBlue, color.FgMagenta, color.FgCyan, color.FgWhite}
	colors := baseColors
	values := []string{"1", "1", "2", "3", "2", "1", "3", "1", "4", "5", "6", "7", "8", "9", "10", "11", "12"}

	sources := make(map[string]color.Attribute)

	for _, v := range values {
		if _, ok := sources[v]; !ok {
			if len(colors) == 0 {
				colors = baseColors
			}
			log.Printf("Adding source %s \n", v)
			sources[v] = colors[0]
			colors = colors[1:]
		} else {
			log.Printf("source %s already added\n", v)
		}
	}

	for _, v := range values {
		if c, ok := sources[v]; ok {
			color.New(c).Println("   -----> " + v)

		}
	}
	/*
		fmt.Println("One LINE")
		// Create a new color object
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Println("Prints cyan text with an underline.")

		// Or just add them to New()
		d := color.New(color.FgCyan, color.Bold)
		d.Printf("This prints bold cyan %s\n", "too!.")
	*/
}
