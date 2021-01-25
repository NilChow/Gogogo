package worker

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func Init() {
	once.Do(initImp)
}

func initImp() {
	fmt.Println("[Worker Init!]")
}