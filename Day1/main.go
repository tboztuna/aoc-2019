package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfRequiredFuel := 0

	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		requiredFuel := math.Floor(mass/3) - 2

		sumOfRequiredFuel += int(requiredFuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The fuel requirements for all of the modules on our spacecraft is: %d", sumOfRequiredFuel)
}
