package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Andreassm99/minyr2/yr"
)

func main() {
	// Wait for the user to enter "minyr"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 'minyr' to start temperature conversion: ")
	text, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(text)) != "minyr" {
		fmt.Println("Invalid input.")
		return
	}

	// Present the user with options
	fmt.Println("Options:")
	fmt.Println("  - 'convert' to convert temperature data from Celsius to Fahrenheit")
	fmt.Println("  - 'average' to get the average temperature for the entire period")
	fmt.Print("Enter option: ")
	option, _ := reader.ReadString('\n')
	option = strings.ToLower(strings.TrimSpace(option))

	if option == "convert" {
		err := yr.Convert()
		if err != nil {
			fmt.Println("Error during temperature conversion:", err)
			return
		}
		fmt.Println("Temperature conversion complete.")
		return
	}

	if option == "average" {
		fmt.Print("Enter unit of measurement ('c' for Celsius or 'f' for Fahrenheit): ")
		unit, _ := reader.ReadString('\n')
		unit = strings.ToLower(strings.TrimSpace(unit))

		avg, err := yr.Average(unit)
		if err != nil {
			fmt.Println("Error calculating average temperature:", err)
			return
		}
		fmt.Printf("Average temperature: %.2f %s\n", avg, unit)
	}

	// Wait for user input
	fmt.Println("Press enter to exit.")
	fmt.Scanln()
}
