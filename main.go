package main

import (
	"github.com/tebeka/selenium"
	"main/components"
)

func main() {
	done := make(chan int)

	Robot := new(components.Robot)
	defer Robot.Service.Stop()
	Robot.Start("https://github.com/ydtg1993")
	Robot.TapIn(selenium.ByXPATH, "//*[@id='js-pjax-container']/div[2]/div/div[2]/div[2]/div/div[1]/div/ol/li[1]/div/div/div/div/a/span")
	<-done
}
