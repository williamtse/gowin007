package win007

import (
	//	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchSchedule(date string) string {
	url := "http://bf.win007.com/football/Next_" + date + ".htm"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	return bodystr
}
