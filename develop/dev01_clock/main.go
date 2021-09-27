package main

import (
	"fmt"
	"log"

	"github.com/makometr/go-tasks-2/develop/dev01_clock/wbtimer"
)

func main() {
	c, err := wbtimer.New(wbtimer.DefaultHost)

	if err != nil {
		log.Fatalln("Error!", err)
	}

	fmt.Println(c.CurrentTime())
}
