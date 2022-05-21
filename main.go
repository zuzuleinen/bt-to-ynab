package main

import (
	"os"

	"github.com/zuzuleinen/bt-to-ynab/converter"
)

func main() {
	argsWithoutProg := os.Args[1:]

	csvPath := argsWithoutProg[0]

	err := converter.Convert(csvPath)
	if err != nil {
		panic(err)
	}
}
