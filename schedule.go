package main

import (
	"fmt"
	//	"./model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"./win007"
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

func ThreadSchedule(ch chan int, scheid string) {
	referee, err := win007.FetchReferee(scheid)
	if err == nil {
		doc, err := goquery.NewDocumentFromResponse(referee)
		if err == nil {
			trs := doc.Find("#content").Eq(0).Find("table").Eq(0).Find("table").Eq(0).Find("tr")

			trs.Each(func(i int, s *goquery.Selection) {

			})
		}
	}
	ch <- 1
}

func main() {
	sches_res, err := win007.FetchSchedule("20170210")
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
			go ThreadSchedule(chs[i], scheid)
		})
		for _, ch := range chs {
			total++
			fmt.Println(strconv.Itoa(total))
			<-ch
		}
	}

	fmt.Println("结束")
}
