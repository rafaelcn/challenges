package main

import (
	"fmt"
	"time"
)

const (
	HHMMSS24h string = "15:04:05"
)

func main() {

	var s, i int64
	var shd, sha string

	fmt.Scanf("%d\n", &s)
	fmt.Scanf("%s\n", &shd)
	fmt.Scanf("%s\n", &sha)

	hd, _ := time.Parse(HHMMSS24h, shd)
	ha, _ := time.Parse(HHMMSS24h, sha)

	hdu := int64(hd.Hour()*60*60+hd.Minute()*60+hd.Second())
	hau := int64(ha.Hour()*60*60+ha.Minute()*60+ha.Second())


	if hdu >= hau {
		hau += 86400 // one day
	}

	for hdu+s <= hau {
		hdu += s
		i++
	}

	fmt.Println(i)
}
