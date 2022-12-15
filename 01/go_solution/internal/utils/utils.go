package utils

import (
	"log"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetInput() []string {
	filename := getFilename()
	data, err := os.ReadFile(filename)
	Check(err)
	return strings.Split(string(data), "\n")
}

func getFilename() string {
	if len(os.Args) < 2 {
		log.Fatal("Must provide a filename!")
	}
	filename := os.Args[1]
	return filename
}
