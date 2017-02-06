package main

import (
	"fmt"

	"./win007"
)

func main() {
	sches_html := win007.FetchSchedule("20170207")
	fmt.Println(sches_html)
}
