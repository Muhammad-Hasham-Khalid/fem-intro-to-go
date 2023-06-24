package main

import "fmt"

func main() {
	var name, city string
	var yearsLived uint
	var isWeatherNice bool

	fmt.Print("What's your name? ")
	fmt.Scan(&name)

	fmt.Print("Which city do you live in? ")
	fmt.Scan(&city)

	fmt.Print("How many years have you lived there? ")
	fmt.Scan(&yearsLived)

	var weatherNice string
	fmt.Print("Is the weather nice in your home town (y/n)? ")
	fmt.Scan(&weatherNice)

	isWeatherNice = weatherNice == "y"

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t\n", name, city, yearsLived, isWeatherNice)
}