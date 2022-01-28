package business

import (
	"fmt"
)

type ioDevice struct{}

func (iod *ioDevice) PrintFinalWinner(text string) {
	fmt.Println(text)
}

var notifier = &ioDevice{}
