package components

import "github.com/tebeka/selenium"

type Robot struct {
	Config    map[string]string
	Service   *selenium.Service
	WebDriver selenium.WebDriver
}
