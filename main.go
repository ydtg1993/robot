package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"main/components"
	"main/tools"
)

func main() {
	done := make(chan int)
	service, wb, err := components.Start(tools.Config())
	defer service.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dom, _ := wb.FindElement(selenium.ByXPATH, "//*[@id='js-pjax-container']/div[2]/div/div[2]/div[2]/div/div[1]/div/ol/li[1]/div/div/div/div/a/span")
	dom.Click()
	<-done
}
