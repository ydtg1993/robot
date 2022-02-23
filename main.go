package main

import (
	"main/components"
)

func main() {
	done := make(chan int)

	Robot := new(components.Robot)
	defer Robot.Service.Stop()
	Robot.Start("https://github.com")
	//Robot.TapIn(selenium.ByXPATH, "//*[@id='__layout']/div/section[1]/div[4]/div/div[1]/ul/li[2]")
	<-done
}
