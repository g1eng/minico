package main

import (
	"log"

	"github.com/g1eng/minico/cmd"
)

func main() {
	if err := cmd.Cmd(); err != nil {
		log.Fatalln(err)
	}
}
