package main

import (
	"fmt"
	//	"./model"
	//	"io"
	"io/ioutil"
	"log"
	"net/http"
	//	"os"
	"regexp"
	"strconv"

	"./src"

	"github.com/PuerkitoBio/goquery"
)

var total int = 0
var db, _ = src.OpenDB()

func ResToStr(response *http.Response) string {
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		return bodystr
	}
	return ""
}

func GetSchedule(ch chan int, s *goquery.Selection, i int) {
	if i > 0 {
		onclick, _ := s.Find("td").Last().Find("a").Eq(0).Attr("onclick")
		leagueTd := s.Find("td").Eq(0)
		color, _ := leagueTd.Attr("bgcolor")
		reg := regexp.MustCompile("[0-9]+")
		scheid := reg.FindString(onclick)

		fmt.Println("赛程id:" + scheid + ";联赛颜色：" + color)
		fmt.Println("采集赛程分析页面" + scheid)
		src.FetchASchedule(scheid, color)
	}

	ch <- 1
}

func main() {
	sches_res, err := src.FetchScheduleFromDate("20170214")
	if err == nil {
		doc, err := goquery.NewDocumentFromResponse(sches_res)
		if err != nil {
			log.Fatal(err)
		}
		tb := doc.Find("#table_live")

		length := tb.Find("tr").Length()
		fmt.Println(length)
		chs := make([]chan int, length)
		tb.Find("tr").Each(func(i int, s *goquery.Selection) {

			chs[i] = make(chan int)
			go GetSchedule(chs[i], s, i)
		})
		for _, ch := range chs {
			total++
			fmt.Println(strconv.Itoa(total))
			<-ch
		}
	}

	fmt.Println("结束")
}
