package service

import (
	"main.go/internal/display"
	"time"

	"github.com/nsf/termbox-go"
)

func Stopwatch() {
	var running bool
	var elapsed time.Duration

	display.Display(elapsed)

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
							display.Display(elapsed)
							time.Sleep(50 * time.Millisecond)
						}
					}()
				}
			}
			if ev.Key == termbox.KeyBackspace {
				return
			}

		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
