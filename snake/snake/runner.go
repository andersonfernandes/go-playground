package snake

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func Run() {
	a := &Arena{
		Snake: []*Snake{
			{
				X:         5,
				Y:         0,
				Direction: Right,
			},
			{
				X:         4,
				Y:         0,
				Direction: Right,
			},
			{
				X:         3,
				Y:         0,
				Direction: Right,
			},
			{
				X:         2,
				Y:         0,
				Direction: Right,
			},
		},
		MaxX: 50,
		MaxY: 25,
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	c := make(chan termbox.Event)
	go listenKeyboard(c)

	go func() {
		for e := range c {
			if e.Key == termbox.KeyCtrlC {
				quit()
			} else {
				a.SetDirection(e.Key)
				go listenKeyboard(c)
			}
		}
	}()

	for {
		a.Move()
		a.Draw()

		time.Sleep(200 * time.Millisecond)
	}
}

func listenKeyboard(c chan termbox.Event) {
	c <- termbox.PollEvent()
}

func quit() {
	termbox.Close()
	os.Exit(1)
}
