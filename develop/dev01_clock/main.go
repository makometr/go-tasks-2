package basedevelop

import (
	"fmt"
	"go-tasks-2/develop/dev01_clock/wbtimer/clock"
	"log"
)

// PrintCurrentTime prints curremt time from remote server
func PrintCurrentTime() {
	// time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	c, err := clock.New(clock.BaseHost)

	if err != nil {
		log.Fatalln("Error!", err)
	}

	fmt.Println("Current time:", c.time)
}
