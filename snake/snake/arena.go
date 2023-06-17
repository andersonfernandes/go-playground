package snake

import (
	"fmt"
	"os"
	"os/exec"
	"time"
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

type Arena struct {
	Snake *SnakePart
	MaxX  int
	MaxY  int
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

			if a.Snake.Contains(Pair{x, y}) {
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
	c := a.Snake
	for c.Next != nil {
		fmt.Println(c.Direction)
		c = c.Next
	}
}

func (a *Arena) HandleUpdate() {
	a.Snake.Move()
	a.Draw()

	time.Sleep(200 * time.Millisecond)
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
