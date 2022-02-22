package components

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"strconv"
)

func Start(config map[string]string) (*selenium.Service, selenium.WebDriver, error) {
	opts := []selenium.ServiceOption{}
	port, _ := strconv.Atoi(config["port"])
	service, err := selenium.NewChromeDriverService(config["selenium"], port, opts...)
	if nil != err {
		return service, nil, errors.New(err.Error())
	}
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
			//"--headless", //
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		},
	}

	caps.AddChrome(chromeCaps)
	wb, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return service, nil, errors.New(err.Error())
	}

	err = wb.Get(config["url"])
	if err != nil {
		return service, nil, errors.New(err.Error())
	}

	return service, wb, nil
}
