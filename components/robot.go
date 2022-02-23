package components

import "github.com/tebeka/selenium"

type Robot struct {
	Config    map[string]string
	Service   *selenium.Service
	WebDriver selenium.WebDriver
}

const USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36"
const SELENIUM_PATH = "chromedriver.exe"
const SELENIUM_PORT = 9510
