package AoC22

import (
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
	return string(dat)
}
