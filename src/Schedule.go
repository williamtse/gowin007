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

	//	"encoding"

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
	ac           *ActiveRecord
}

//采集单个赛程
func FetchASchedule(scheid string, leagueColor string) {
	if scheid == "0" {
		return
	}
	fmt.Println("采集分析页面" + scheid)
	resp, err := FetchAnanly(scheid)
	if err == nil {
		defer resp.Body.Close()
	}

	if err == nil {
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err == nil {
			schetr := doc.Find(".t1p1").Eq(0).Find("table").Eq(0).Find("tr").Eq(0)
			tds := schetr.Children()
			himg, _ := tds.Eq(0).Find("img").Eq(0).Attr("src")
			cimg, _ := tds.Eq(2).Find("img").Eq(0).Attr("src")
			fmt.Println(himg)
			infotd := tds.Eq(1)
			infoTds := infotd.Find("td")
			teamLinks := infoTds.Eq(0).Find("a")
			fmt.Println(teamLinks)
			hTeamLink := teamLinks.Eq(0)
			cTeamLink := teamLinks.Eq(1)
			hTeamHref, _ := hTeamLink.Attr("href")
			cTeamHref, _ := cTeamLink.Attr("href")
			fmt.Println(hTeamHref)
			fmt.Println(cTeamHref)
			hTeamName := GetName(hTeamLink.Text())
			cTeamName := GetName(cTeamLink.Text())
			fmt.Println(hTeamName)
			fmt.Println(cTeamName)

			reg := regexp.MustCompile("http://info\\.win007\\.com/team/([0-9]+)")
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
			leagueIdMatches := GetMatches("([0-9]+)\\.htm", leagueHref)

			leagueId := leagueIdMatches[1]
			leagueNameMatches := strings.Split(leagueStr, " ")

			leagueName := leagueNameMatches[0]
			ar := ActiveRecord{table: "league"}
			league, err := ar.Find("name='" + leagueName + "'")
			if !league {
				ar.isNew = true
				league = League{name: leagueName, color: color, ar: ar}
				league.Save()
			}

			fmt.Println(leagueId)
			fmt.Println(leagueName)
			if len(time) < 5 {
				time = "0" + time
			}
			startTime := date + " " + time
			startTimeUnix := TimeToUnix(startTime)
			fmt.Println(startTimeUnix)

		} else {
			fmt.Println("采集分析页面失败：" + err.Error())
		}
	}

}

func (sche *Schedule) save() error {
	db, err := OpenDB()
	CheckErr(err)
	rows, err := db.Query("select * from " + sche.Table + " where id=" + scheid)
	CheckErr(err)
	fmt.Println(rows)
	if !sche.ar.isNew {
		stmt, err := db.Prepare("INSERT INTO " + sche.ar.table + " (id,)values(?,?)")
	} else {
		stmt, err := db.Prepare("UPDATE " + sche.ar.table + " set ")
	}
	CheckErr(err)
	res, err := stmt.Exec(scheid, sche)
	CheckErr(err)
	fmt.Println(res)
	defer db.Close()
	return res, err
}
