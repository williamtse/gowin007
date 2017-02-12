package src

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	//	"regexp"
)

type Team struct {
	Id   int
	name string
}

func getImg(teamId string, url string) (n int64, err error) {
	fileName := "./imgs/" + teamId + ".png"
	os.MkdirAll(path.Dir(fileName), os.ModePerm)
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件创建错误" + err.Error())

	}
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}
func GetName(str string) (name string) {
	matches := GetMatches("([^\\(]*).*", str)
	rename := matches[1]
	return rename
}
