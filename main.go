package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func main() {
	var g g2048

	err := termbox.Init()
	if err != nil {
		fmt.Printf("初始化失败, Error:%s\n", err)
		os.Exit(1)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	g.Run()
}
