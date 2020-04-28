package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

func main() {
	pin := rpio.Pin(14)
	if err := rpio.Open(); err != nil {
		fmt.Printf("cannot open GPIO pin")
	}
	pin.Input()
	pin.PullDown()
	pin.Detect(rpio.AnyEdge)
	fmt.Printf("press...")
	for {
		sure := 0
		for {
			if sure >= 2 {
				break
			}
			if pin.EdgeDetected() {
				fmt.Printf("pressed\n")
				sure++
			}
			time.Sleep(time.Second / 2)
		}
		fmt.Println("cycle ended")
		time.Sleep(time.Second * 2)

	}
}
