package display

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func Display(elapsed time.Duration) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	timeStr := fmt.Sprintf("Time: %.5f seconds", elapsed.Seconds())

	for i, ch := range timeStr {
		termbox.SetCell(i, 1, ch, termbox.ColorBlack, termbox.ColorGreen)
	}

	termbox.Flush()
}
