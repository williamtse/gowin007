package src

import (
	"fmt"
	"strings"
	//	"io"
	//	"io/ioutil"
	//	"log"
	//	"net/http"
	//	"os"
	"regexp"
	//	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Schedule struct {
	Odds         []Odd
	OriginOdds   []Odd
	Hteam        Team
	Cteam        Team
	League       League
	StartTime    int
	SecStartTime int
	Status       int
	Scores       Score
}

//采集单个赛程
func FetchASchedule(scheid string) {
	fmt.Println("采集教练页面" + scheid)
	analy, err := FetchAnanly(scheid)
	if err == nil {
		doc, err := goquery.NewDocumentFromResponse(analy)
		if err == nil {
			schetr := doc.Find(".t1p1").Eq(0).Find("table").Eq(0).Find("tr").Eq(0)
			tds := schetr.Children()
			himg, _ := tds.Eq(0).Find("img").Eq(0).Attr("src")
			cimg, _ := tds.Eq(2).Find("img").Eq(0).Attr("src")

			infotd := tds.Eq(1)
			infoTds := infotd.Find("td")
			teamLinks := infoTds.Eq(0).Find("a")

			hTeamLink := teamLinks.Eq(0)
			cTeamLink := teamLinks.Eq(1)
			hTeamHref, _ := hTeamLink.Attr("href")
			cTeamHref, _ := cTeamLink.Attr("href")
			//			hTeamName := GetName(hTeamLink.Text())
			//			cTeamName := GetName(cTeamLink.Text())

			reg := regexp.MustCompile("http://info\\.win007\\.com/team/([\\d]+)\\.htm")
			hteamIdMatch := reg.FindStringSubmatch(hTeamHref)
			cteamIdMatch := reg.FindStringSubmatch(cTeamHref)
			hteamId := hteamIdMatch[1]
			cteamId := cteamIdMatch[1]
			getImg(hteamId, himg)
			getImg(cteamId, cimg)

			timeTd := infoTds.Eq(1)
			timeStr := timeTd.Text()
			fmt.Println(timeStr)
			timeMatches := GetMatches("[^\\d]*([\\d]{4}-[\\d]{2}-[\\d]{2})[^\\d]*([\\d]{1,2}:[\\d]{2})[^\\d]*", timeStr)
			date := timeMatches[1]
			time := timeMatches[2]

			leagueLink := infoTds.Eq(2).Find("a").Eq(0)
			leagueHref, _ := leagueLink.Attr("href")
			leagueStr := leagueLink.Text()
			leagueIdMatches := GetMatches("http://info\\.win007\\.com/league_match/league_vs/[0-9]*\\-*[0-9]*/([0-9]+)", leagueHref)
			leagueId := leagueIdMatches[1]
			leagueNameMatches := strings.Split(leagueStr, " ")

			leagueName := leagueNameMatches[0]
			fmt.Println(leagueName)
			if len(time) < 5 {
				time = "0" + time
			}
			startTime := date + " " + time
			startTimeUnix := TimeToUnix(startTime)
			fmt.Println(startTimeUnix)
			tds.Each(func(i int, s *goquery.Selection) {

			})
		} else {
			fmt.Println("采集教练页面失败")
		}
	}

}

func save(ac *ActiveRecord, sche *Schedule) Error {
	ac.tableName("schedule")

}
