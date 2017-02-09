package win007

import (
	//	"fmt"
	//	"io/ioutil"
	"log"
	"net/http"
)

func FetchSchedule(date string) (*http.Response, error) {
	url := "http://bf.win007.com/football/Next_" + date + ".htm"
	return SimpleGet(url)
}

func FetchAnaly(scheid string) (*http.Response, error) {
	url := "http://zq.win007.com/analysis/" + scheid + "cn.htm"
	return SimpleGet(url)
}

func FetchReferee(scheid string) (*http.Response, error) {
	url := "http://zq.win007.com/referee/" + scheid + "cn.htm"
	return SimpleGet(url)
}

func SimpleGet(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.Do(req)
	return resp, err
}
