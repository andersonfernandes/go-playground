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

		s := a.Snake
		for s.Next != nil {
			fmt.Println(s.Moves)
			s = s.Next
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

func drawXBorder(max int, rCorner string, lCorner string) {
	for i := 0; i < max; i++ {
		if i == 0 {
			fmt.Print(rCorner)
		}

		if i < max-1 {
			fmt.Print(XBorder)
		} else {
			fmt.Println(XBorder + lCorner)
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
	s := a.Snake
	for s.Next != nil {
		o := "- "
		for i, m := range s.Moves {
			o += m

			if i != len(s.Moves)-1 {
				o += ", "
			}
		}

		f.Write([]byte(o + "\n"))
		s = s.Next
	}

	f.Close()
}
