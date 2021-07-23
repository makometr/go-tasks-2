package basedevelop

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

// PrintCurrentTime prints curremt time from remote server
func PrintCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln("Error!", err)
	}

	fmt.Println("Current time:", time)
}
