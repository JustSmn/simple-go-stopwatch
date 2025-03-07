package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	//--------------------------- Инициализация termbox
	err := termbox.Init()
	if err != nil {
		fmt.Println("Error initializing termbox:", err)
		return
	}
	defer termbox.Close()
	//--------------------------------------------------

	var running bool
	var elapsed time.Duration

	display(elapsed)

	for {

		ev := termbox.PollEvent()

		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeySpace {
				if running {
					running = false
				} else {
					running = true
					startTime := time.Now()

					go func() {
						for running {
							elapsed = time.Since(startTime)
							display(elapsed)
							time.Sleep(50 * time.Millisecond)
						}
					}()
				}
			}
			if ev.Key == termbox.KeyEsc {
				return
			}

		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func display(elapsed time.Duration) {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	timeStr := fmt.Sprintf("Time: %.5f seconds", elapsed.Seconds())

	for i, ch := range timeStr {
		termbox.SetCell(i, 1, ch, termbox.ColorBlack, termbox.ColorGreen)
	}

	termbox.Flush()
}
