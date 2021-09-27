package basedevelop

import (
	"log"

	"github.com/makometr/go-tasks-2/develop/dev01_clock/wbtimer"
)

// PrintCurrentTime prints curremt time from remote server
func PrintCurrentTime() {
	var c wbtimer.Clock
	// c, err := clock.New(clock.BaseHost)
	// wbtimer.Clock
	var err error

	if err != nil {
		log.Fatalln("Error!", err)
	}

	// fmt.Println("Current time:", c.CurrentTime())
}
