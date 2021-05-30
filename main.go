package main

import (
	"github.com/g1eng/minico/v2/cmd"
	"log"
)

func main() {
	if err := cmd.Cmd(); err != nil {
		log.Fatalln(err)
	}
}
