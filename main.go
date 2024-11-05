package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Temperature converter")
	fmt.Println(
		"Choose a Kelvin temperature that you want to convert. (e.g: Water boiling point = 373.2 K)",
	)
	var k float64
	_, err := fmt.Scan(&k)
	if err != nil {
		log.Fatal(err)
	}
	c := k - 273.16
	fmt.Printf("Celsius Temperature: %f°C\n", c)

}
