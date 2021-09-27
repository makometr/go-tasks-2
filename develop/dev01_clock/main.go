package basedevelop

import (
	"fmt"
	"log"

	"github.com/makometr/go-tasks-2/develop/dev01_clock/wbtimer"
)

// PrintCurrentTime prints curremt time from remote server
func PrintCurrentTime() {
	c, err := wbtimer.New(wbtimer.BaseHost)

	if err != nil {
		log.Fatalln("Error!", err)
	}

	fmt.Println(c.CurrentTime())
}
