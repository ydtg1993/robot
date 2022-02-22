package components

import (
	"fmt"
	"strconv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func Start(config map[string]string)  {
	opts := []selenium.ServiceOption{}
	port,_ := strconv.Atoi(config["port"])
	service, err := selenium.NewChromeDriverService(config["selenium"], port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
	}
	defer service.Stop()
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
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		},
	}

	caps.AddChrome(chromeCaps)
	wb, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	//但是不会导致seleniumServer关闭
	defer wb.Quit()
	err = wb.Get(config["url"])
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
}
