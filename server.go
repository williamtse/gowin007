package main

import (
	"fmt"

	"./model"

	"./win007"
)

func main() {
	sches_html := win007.FetchSchedule("20170207")
	//	fmt.Println(sches_html)
	a_odd := model.Odd{}
	a_odd.H = 1.12
	a_odd.P = 2.5
	a_odd.G = 3.5
	schedule := model.Schedule{}
	schedule.A_odd = a_odd
	fmt.Println(schedule)
}
