package snake

import (
	"os"

	"github.com/nsf/termbox-go"
)

func Run() {
	m := ModeNormal
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		m = ModeDebug
	}

	a := &Arena{
		Snake: &SnakePart{
			Position: Pair{5, 0},
			Moves:    []string{Right},
			Next: &SnakePart{
				Position: Pair{4, 0},
				Moves:    []string{Right},
				Next: &SnakePart{
					Position: Pair{3, 0},
					Moves:    []string{Right},
					Next: &SnakePart{
						Position: Pair{2, 0},
						Moves:    []string{Right},
					},
				},
			},
		},
		Mode: m,
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
				if e.Key == termbox.KeyArrowUp || e.Key == termbox.KeyArrowDown || e.Key == termbox.KeyArrowRight || e.Key == termbox.KeyArrowLeft {
					a.Snake.UpdateDirection(e.Key)
				}

				go listenKeyboard(c)
			}
		}
	}()

	for {
		a.HandleUpdate()
	}
}

func listenKeyboard(c chan termbox.Event) {
	c <- termbox.PollEvent()
}

func quit() {
	termbox.Close()
	os.Exit(0)
}
