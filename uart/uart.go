package uart

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Uart struct {
	Rx rpio.Pin
	Tx rpio.Pin
}

var (
	ticker = time.NewTicker(time.Second)
)

func (uart Uart) Read() (buf uint8) {
	uart.Rx.Input()
	var count = 0
	for range ticker.C {
		if uart.Rx.Read() == rpio.Low {
			break
		}
	}
	for t := range ticker.C {

		if count < 8  {
			buf = buf >> 1
			bit := uint8(uart.Rx.Read())
			fmt.Println(count, t, bit)
			//buf = buf | (bit << 7)
			if bit == 0x01 {
				buf = buf | 0x80
			}
			count++
		} else {
			break
		}
	}
	return
}
