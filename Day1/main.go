package main

import (
	"bufio"
	"log"
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

	for scanner.Scan() {
		mass := strconv.Atoi(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
