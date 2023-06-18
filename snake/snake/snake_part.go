package snake

import "github.com/nsf/termbox-go"

type Pair [2]int

type SnakePart struct {
	Position Pair
	Moves    []string
	Next     *SnakePart
}

func (s *SnakePart) Contains(c Pair) bool {
	for s.Next != nil {
		if s.Position == c {
			return true
		}

		s = s.Next
	}

	return false
}

func (s *SnakePart) Eat() {
	if s == nil {
		s = &SnakePart{Position: Pair{0, 0}, Moves: []string{Right}}
		return
	}

	for s.Next != nil {
		s = s.Next
	}

	d := s.Moves[len(s.Moves)-1]
	ns := &SnakePart{Moves: []string{d}}
	switch d {
	case Up:
		ns.Position = Pair{s.Position[0], s.Position[1] + 1}
	case Down:
		ns.Position = Pair{s.Position[0], s.Position[1] - 1}
	case Right:
		ns.Position = Pair{s.Position[0] - 1, s.Position[1]}
	case Left:
		ns.Position = Pair{s.Position[0] + 1, s.Position[1]}
	}

	s.Next = ns
}

func (s *SnakePart) Move() {
	var d string

	for s.Next != nil {
		d, s.Moves = s.Moves[0], s.Moves[1:]

		switch d {
		case Up:
			s.Position[1] -= 1
		case Down:
			s.Position[1] += 1
		case Right:
			s.Position[0] += 1
		case Left:
			s.Position[0] -= 1
		}

		if len(s.Moves) == 0 {
			s.Moves = append(s.Moves, d)
		}

		s = s.Next
	}
}

func (s *SnakePart) UpdateDirection(k termbox.Key) {
	d := directionFromKey(k)
	s.Moves = []string{d}

	cs := s.Next
	pm := s.Moves[len(s.Moves)-1]
	for cs.Next != nil {
		tm := cs.Moves[len(cs.Moves)-1]

		cs.Moves = append(cs.Moves, pm)
		cs.Moves = append(cs.Moves, d)

		pm = tm
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
