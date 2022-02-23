package components

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"main/tools"
	"strconv"
)

func (Robot *Robot) Start(url string) {
	Robot.Config = tools.Config()
	opts := []selenium.ServiceOption{}
	port, _ := strconv.Atoi(Robot.Config["port"])
	service, err := selenium.NewChromeDriverService(Robot.Config["selenium"], port, opts...)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	Robot.Service = service
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless",
			//"--no-sandbox",
			USER_AGENT,
		},
	}

	caps.AddChrome(chromeCaps)
	wb, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = wb.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Robot.WebDriver = wb
	return
}
