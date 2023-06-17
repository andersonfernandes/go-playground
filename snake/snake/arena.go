package snake

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	ModeNormal string = "NORMAL"
	ModeDebug         = "DEBUG"
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
	Mode  string
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

	if a.Mode == ModeDebug {
		fmt.Println("[DEBUG]", time.Now())

		c := a.Snake
		for c.Next != nil {
			fmt.Println(c.Direction)
			c = c.Next
		}
	}

	writeLog(*a)
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

func writeLog(a Arena) {
	if a.Mode != ModeDebug {
		return
	}

	f, err := os.OpenFile(
		"snake_run.log",
		os.O_RDWR|os.O_APPEND|os.O_CREATE,
		0644,
	)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	f.Write([]byte("\n---\n"))
	cs := a.Snake
	for cs.Next != nil {
		o := "- "
		for i, s := range cs.Direction {
			o += s

			if i != len(cs.Direction)-1 {
				o += ", "
			}
		}

		f.Write([]byte(o + "\n"))
		cs = cs.Next
	}

	f.Close()
}
