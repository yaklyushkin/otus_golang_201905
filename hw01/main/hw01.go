package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	ntpT, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	t := time.Now()
	fmt.Println("Текущее время:\t\t" + t.Format(time.RFC3339Nano))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Точно время NTP:\t" + ntpT.Format(time.RFC3339Nano))
	}
}
