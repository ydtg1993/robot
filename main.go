package main

import (
	"main/components"
)

func main() {
	done := make(chan int)
	Robot := new(components.Robot)
	Robot.Start()
	<-done
}
