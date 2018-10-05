package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"helloworld/uart"
	"time"
)

func main() {
	rpio.Open()
	defer rpio.Close()

	go read()
	go send()

	time.Sleep(time.Second * 15)
}

func read() {
	u := uart.Uart{
		Rx: rpio.Pin(4),
		Tx: rpio.Pin(5),
	}

	var res = u.Read()
	fmt.Println(res)
}

func send() {
	cnt := 0
	p3 := rpio.Pin(17)

	p3.Output()
	p3.High()
	time.Sleep(time.Second * 3)

	//
	var number uint8
	number = 53

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		if cnt == 0 {
			cnt++
			p3.Low()
		} else {
			cnt++
			var bit8 uint8
			bit8 = number << 7
			p3.Write(rpio.State(bit8 >> 7))
			if cnt == 9 {
				break
			}
			number = number >> 1
		}

	}
}
