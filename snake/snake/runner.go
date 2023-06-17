package snake

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func Run() {
	a := &Arena{
		SnakeHead: &SnakePart{
			Position:  Pair{5, 0},
			Direction: []string{Right},
			Next: &SnakePart{
				Position:  Pair{4, 0},
				Direction: []string{Right},
				Next: &SnakePart{
					Position:  Pair{3, 0},
					Direction: []string{Right},
					Next: &SnakePart{
						Position:  Pair{2, 0},
						Direction: []string{Right},
					},
				},
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
