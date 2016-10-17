package main

import (
	"fmt"
	"time"
)

func main() {
	const layout = "2006-01-02"

	base := time.Date(2016, 11, 1, 0, 0, 0, 0, time.Local)

	for index := 2; index < 8; index++ {
		estimate := base.AddDate(0, index, -1)
		fmt.Println(estimate.Format(layout))
	}
}
