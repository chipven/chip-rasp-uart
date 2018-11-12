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
	ticker = time.NewTicker(time.Second / 900)
)

func (uart Uart) Read32() (buf uint32) {
	uart.Rx.Input()
	var count = 0
	for range ticker.C {

		if count == 0 || count == 10 || count == 20 || count == 30 {
			if uart.Rx.Read() == rpio.Low {
				count++
				continue
			}
		}

		if count == 9 || count == 19 || count == 29 {
			count++
		}
		if (count > 0 && count < 9) ||
			(count > 10 && count < 19) ||
			(count > 20 && count < 29) ||
			(count > 30 && count < 39) {
			buf = buf >> 1
			bit := uint8(uart.Rx.Read())
			fmt.Print(bit) //fmt包下的Println()方法
			//fmt.Println(count, t, bit)
			//buf = buf | (bit << 7)
			if bit == 0x01 {
				buf = buf | 0x80000000
			}
			count++
		}
		if count > 38 {
			break
		}

	}
	return

}

func (uart Uart) Read() (buf uint8) {
	uart.Rx.Input()
	var count = 0
	for range ticker.C {
		if uart.Rx.Read() == rpio.Low {
			break
		}
	}
	for range ticker.C {

		if count < 8 {
			buf = buf >> 1
			bit := uint8(uart.Rx.Read())
			//fmt.Println(count, t, bit)
			//buf = buf | (bit << 7)
			if bit == 0x01 {
				buf = buf | 0x80
			}
			count++
		} else {
			break
		}
	}
	<-ticker.C
	return
}
