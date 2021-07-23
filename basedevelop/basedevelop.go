package basedevelop

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// PrintCurrentTime prints curremt time from remote server
func PrintCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error! ", err)
	}

	fmt.Println("Current time:", time)
}
