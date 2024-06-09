package service1

import (
	"fmt"

	"github.com/ceocoder/dexter"
)

func Run() {
	fmt.Println("running service1")
	dex := dexter.NewDexter()

	fmt.Println("wait and kill 3")
	dex.WaitAndKill()
}
