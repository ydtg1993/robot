package components

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (Robot *Robot) TapIn(method, path string) bool {
	for i := 0; i < 3; i++ {
		dom, err := Robot.WebDriver.FindElement(method, path)
		if err != nil {
			continue
		}
		dom.Click()
		time.Sleep(1 * time.Second)
		//检查dom 判断跳转
		_, err = Robot.WebDriver.FindElement(method, path)
		if err == nil {
			return true
		}
	}
	return false
}

func (Robot *Robot) CatchApi(url string, method string, body io.Reader) []byte {
	payload := strings.NewReader("page=1&limit=500&start=2012-11-24%2000%3A00%3A00&end=2019-11-24%2023%3A00%3A00&pid=&rid=")
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("Accept", "application/json, text/javascript, */*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	current_url, _ := Robot.WebDriver.CurrentURL()
	req.Header.Add("Referer", current_url)
	req.Header.Add("Origin", current_url)
	req.Header.Add("Host", current_url)
	req.Header.Add("User-Agent", USER_AGENT)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()

	var response []byte
	if err != nil {
		fmt.Println(err.Error())
		return response
	}

	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return response
	}
	response, _ = ioutil.ReadAll(reader)

	return response
}
