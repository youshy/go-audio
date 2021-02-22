package main

import (
	"github.com/youshy/go-audio/sine"
)

func main() {
	generateSineToFile()
}

func generateSineToFile() {
	err := sine.GenerateTunedSineToFile(523.25, "out.bin")
	if err != nil {
		panic(err)
	}
}
