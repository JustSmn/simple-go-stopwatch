package main

import (
	"fmt"
	"main.go/internal/service"
	termbox "main.go/internal/tbox"
)

func main() {
	closeTermbox, err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeTermbox()

	service.Stopwatch()
}
