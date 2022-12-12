package AoC22

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads file from path provided.
func ReadInput(fileLocation string) string {
	dat, err := os.ReadFile(fileLocation)
	check(err)
	fmt.Println(string(dat))

	return string(dat)
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
