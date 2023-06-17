package snake

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	Up    string = "TOP"
	Down         = "DOWN"
	Right        = "RIGHT"
	Left         = "LEFT"
)

const (
	XBorder string = "───"
	YBorder        = "│"
)

type Pair [2]int

type SnakePart struct {
	Position  Pair
	Direction []string
	Next      *SnakePart
}

type Arena struct {
	SnakeHead *SnakePart
	MaxX      int
	MaxY      int
}

func (a *Arena) eat() {
	if a.SnakeHead == nil {
		a.SnakeHead = &SnakePart{Position: Pair{0, 0}, Direction: []string{Right}}
		return
	}

	cs := a.SnakeHead
	for cs.Next != nil {
		cs = cs.Next
	}

	d := cs.Direction[len(cs.Direction)-1]
	s := &SnakePart{Direction: []string{d}}
	switch d {
	case Up:
		s.Position = Pair{cs.Position[0], cs.Position[1] + 1}
	case Down:
		s.Position = Pair{cs.Position[0], cs.Position[1] - 1}
	case Right:
		s.Position = Pair{cs.Position[0] - 1, cs.Position[1]}
	case Left:
		s.Position = Pair{cs.Position[0] + 1, cs.Position[1]}
	}

	cs.Next = s
}

func (a *Arena) Move() {
	cs := a.SnakeHead
	var d string

	for cs.Next != nil {
		d, cs.Direction = cs.Direction[0], cs.Direction[1:]

		switch d {
		case Up:
			cs.Position[1] -= 1
		case Down:
			cs.Position[1] += 1
		case Right:
			cs.Position[0] += 1
		case Left:
			cs.Position[0] -= 1
		}

		if len(cs.Direction) == 0 {
			cs.Direction = append(cs.Direction, d)
		}

		cs = cs.Next
	}
}

func (a *Arena) SetDirection(k termbox.Key) {
	d := directionFromKey(k)
	a.SnakeHead.Direction = []string{d}

	cs := a.SnakeHead.Next
	p := a.SnakeHead.Direction[len(a.SnakeHead.Direction)-1]
	for cs.Next != nil {
		tp := cs.Direction[len(cs.Direction)-1]

		cs.Direction = append(cs.Direction, p)
		cs.Direction = append(cs.Direction, d)

		p = tp
		cs = cs.Next
	}
}

func directionFromKey(k termbox.Key) string {
	var d string

	switch k {
	case termbox.KeyArrowDown:
		d = Down
	case termbox.KeyArrowUp:
		d = Up
	case termbox.KeyArrowRight:
		d = Right
	case termbox.KeyArrowLeft:
		d = Left
	}

	return d
}

func contains(sh *SnakePart, c Pair) bool {
	for sh.Next != nil {
		if sh.Position == c {
			return true
		}

		sh = sh.Next
	}

	return false
}

func (a *Arena) Draw() {
	clear()

	drawLogo()
	drawXBorder(a.MaxX, "┌", "┐")
	for y := 0; y < a.MaxY; y++ {
		for x := 0; x < a.MaxX; x++ {
			if x == 0 {
				fmt.Print(YBorder)
			}

			if contains(a.SnakeHead, Pair{x, y}) {
				fmt.Print("■  ")
			} else {
				fmt.Print("   ")
			}

			if x == a.MaxX-1 {
				fmt.Println(YBorder)
			}
		}
	}
	drawXBorder(a.MaxX, "└", "┘")
	fmt.Println("-> ", time.Now())

	c := a.SnakeHead
	for c.Next != nil {
		fmt.Println(c.Direction)
		c = c.Next
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func drawLogo() {
	fmt.Println("                                                                                              ")
	fmt.Println("                                                                                              ")
	fmt.Println("       ▄████████ ███▄▄▄▄      ▄████████    ▄█   ▄█▄    ▄████████         ▄██████▄   ▄██████▄  ")
	fmt.Println("      ███    ███ ███▀▀▀██▄   ███    ███   ███ ▄███▀   ███    ███        ███    ███ ███    ███ ")
	fmt.Println("      ███    █▀  ███   ███   ███    ███   ███▐██▀     ███    █▀         ███    █▀  ███    ███ ")
	fmt.Println("      ███        ███   ███   ███    ███  ▄█████▀     ▄███▄▄▄           ▄███        ███    ███ ")
	fmt.Println("    ▀███████████ ███   ███ ▀███████████ ▀▀█████▄    ▀▀███▀▀▀          ▀▀███ ████▄  ███    ███ ")
	fmt.Println("             ███ ███   ███   ███    ███   ███▐██▄     ███    █▄         ███    ███ ███    ███ ")
	fmt.Println("       ▄█    ███ ███   ███   ███    ███   ███ ▀███▄   ███    ███        ███    ███ ███    ███ ")
	fmt.Println("     ▄████████▀   ▀█   █▀    ███    █▀    ███   ▀█▀   ██████████        ████████▀   ▀██████▀  ")
	fmt.Println("                                                                                              ")
	fmt.Println("                                                                                              ")
}

func drawXBorder(max int, rc string, lc string) {
	for x := 0; x < max; x++ {
		if x == 0 {
			fmt.Print(rc)
		}

		if x < max-1 {
			fmt.Print(XBorder)
		} else {
			fmt.Println(XBorder + lc)
		}
	}
}
