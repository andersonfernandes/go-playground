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

type Snake struct {
	X         int
	Y         int
	Direction string
	ChangeIn  int
	ChangeTo  string
}

type Arena struct {
	Snake []*Snake
	MaxX  int
	MaxY  int
}

func (a *Arena) Move() {
	for _, s := range a.Snake {
		switch s.Direction {
		case Up:
			s.Y -= 1
		case Down:
			s.Y += 1
		case Right:
			s.X += 1
		case Left:
			s.X -= 1
		}

		if s.ChangeIn == 1 {
			s.Direction = s.ChangeTo
		}
		s.ChangeIn -= 1
	}
}

func (a *Arena) SetDirection(k termbox.Key) {
	d := directionFromKey(k)
	ci := 0

	for i, s := range a.Snake {
		if i == 0 {
			s.Direction = d
			s.ChangeTo = d
		} else {
			s.ChangeTo = d
		}

		s.ChangeIn = ci
		ci += 1
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

func contains(sl []*Snake, c [2]int) bool {
	for _, s := range sl {
		if [2]int{s.X, s.Y} == c {
			return true
		}
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

			if contains(a.Snake, [2]int{x, y}) {
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
