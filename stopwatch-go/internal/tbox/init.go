package termbox

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func Init() (func(), error) {
	err := termbox.Init()
	if err != nil {
		return nil, fmt.Errorf("error initializing termbox: %w", err)
	}
	return termbox.Close, nil
}
