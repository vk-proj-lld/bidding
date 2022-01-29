package uc

import (
	"fmt"
	"sync"

	"github.com/vkstack/bidding/interfaces/notifiers"
)

type outUseCase struct {
	mu sync.Mutex
}

func NewConsoleOutPutUsecase() notifiers.IOut {
	return &outUseCase{}
}

func (out *outUseCase) Write(contents ...interface{}) {
	out.mu.Lock()
	defer out.mu.Unlock()
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	for content := range contents {
		fmt.Println(content)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
}
