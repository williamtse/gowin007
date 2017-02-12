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

func ResToStr(response *http.Response) string {
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		return bodystr
	}
	return ""
}

func GetSchedule(ch chan int, scheid string) {
	fmt.Println("开启采集赛程线程" + scheid)
	src.FetchASchedule(scheid)
	ch <- 1
}

func main() {
	sches_res, err := src.FetchScheduleFromDate("20170210")
	if err == nil {
		doc, err := goquery.NewDocumentFromResponse(sches_res)
		if err != nil {
			log.Fatal(err)
		}
		tb := doc.Find("#table_live")

		length := tb.Find("tr").Length()
		fmt.Println(length)
		chs := make([]chan int, length)
		scheid := "0"
		tb.Find("tr").Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				onclick, _ := s.Find("td").Last().Find("a").Eq(0).Attr("onclick")
				reg := regexp.MustCompile("[0-9]+")
				scheid = reg.FindString(onclick)
			}
			chs[i] = make(chan int)
			go GetSchedule(chs[i], scheid)
		})
		for _, ch := range chs {
			total++
			fmt.Println(strconv.Itoa(total))
			<-ch
		}
	}

	fmt.Println("结束")
}
