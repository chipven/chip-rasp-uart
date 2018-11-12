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

	for {
		time.Sleep(time.Second * 15)
	}
}

func read() {

	u := uart.Uart{
		Rx: rpio.Pin(4),
		Tx: rpio.Pin(5),
	}

	for {
		u32 := u.Read32()
		fmt.Println(u32)
	}

	//cnt := 0
	//for i := 0; i < 500; i++ {
	//	var res = u.Read()
	//	fmt.Println(i, res)
	//if res != 97 {
	//	cnt++
	//	break
	//}
	//}
}

//
//func send() {
//	cnt := 0
//	p3 := rpio.Pin(17)
//
//	p3.Output()
//	p3.High()
//	time.Sleep(time.Second * 3)
//
//	//
//	var number uint8
//	number = 53
//
//	ticker := time.NewTicker(time.Second / 900)
//	for range ticker.C {
//		if cnt == 0 {
//			cnt++
//			p3.Low()
//		} else {
//			cnt++
//			var bit8 uint8
//			bit8 = number << 7
//			p3.Write(rpio.State(bit8 >> 7))
//			if cnt == 9 {
//				break
//			}
//			number = number >> 1
//		}
//
//	}
//}
