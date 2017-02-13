package src

import (
	"fmt"
	//	"io/ioutil"
	//	"log"
	"net/http"
	"time"
)

func FetchScheduleFromDate(date string) (*http.Response, error) {
	url := "http://bf.win007.com/football/Next_" + date + ".htm"
	return SimpleGet(url)
}

func FetchAnaly(scheid string) (*http.Response, error) {
	url := "http://m.win007.com/Analy/Analysis.aspx?scheid=" + scheid
	return SimpleGet(url)
}

func FetchReferee(scheid string) (*http.Response, error) {
	url := "http://zq.win007.com/referee/" + scheid + "cn.html"
	return SimpleGet(url)
}

func FetchAnanly(scheid string) (*http.Response, error) {
	url := "http://zq.win007.com/analysis/" + scheid + "sb.htm"
	fmt.Println(url)
	return SimpleGet(url)
}

func SimpleGet(url string) (*http.Response, error) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	return resp, err
}
