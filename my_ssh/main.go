package main

import (
	"log"
	"my_ssh/tools"
)

func main() {
	if err := tools.Script(); err != nil {
		log.Fatal(err)
	}
}
