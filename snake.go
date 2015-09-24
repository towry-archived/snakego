package main

import "time"
import tl "github.com/JoelOtter/termloop"

type Direct int 

const (
	KeyArrowUp Direct = -1
	KeyArrowRight  Direct = -2
	KeyArrowLeft  Direct = 2
	KeyArrowDown  Direct = 1
)

type Snake struct {
	head *tl.Entity
	body []*Body
	food *Food
	level *tl.BaseLevel
	px int 
	py int
	dir Direct
	size int
	stop bool
	update time.Time
}

func NewSnake(game *tl.Game) (*Snake) {	

	snake := new(Snake)
	snake.head = tl.NewEntity(1, 1, 1, 1)
	snake.px = 1
	snake.py = 0
	snake.head.SetPosition(snake.px, snake.py)
	snake.head.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '#'})
	snake.dir = KeyArrowRight
	snake.stop = false
	snake.update = time.Now()

	snake.level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorCyan,
	})

	snake.level.AddEntity(snake)

	snake.body = make([]*Body, 10)
	snake.size = 0

	// create food
	snake.food = NewFood()
	snake.level.AddEntity(snake.food)

	return snake
}

func (s *Snake) Size() (int, int) {
	return s.head.Size()
}

func (s *Snake) Position() (int, int) {
	return s.head.Position()
}

func (s *Snake) Collide(collision tl.Physical) {
	if _, ok := collision.(*Food); ok {
		if cap(s.body) == s.size {
			rooms := make([]*Body, s.size * 2)
			copy(rooms, s.body)
			s.body = rooms
		}
		var x, y int
		var d Direct
		if s.size == 0 {
			x, y = s.head.Position()
			d = s.dir
		} else {
			x, y = s.body[s.size - 1].Position()
			d = s.body[s.size - 1].Direct()
		}
		switch d {
			case KeyArrowRight:
				x -= 1
				break
			case KeyArrowLeft:
				x += 1
				break
			case KeyArrowUp: 
				y += 1
				break
			case KeyArrowDown:
				y -= 1
				break
		}

		s.body[s.size] = NewBody(x, y, d)
		s.size += 1
		s.food.Invalid()
	}
}

func (s *Snake) Draw(screen *tl.Screen) {
	// check border collision
	w, h := screen.Size()

	if s.px > w {
		s.px = 0
	} else if s.px < 0 {
		s.px = w
	}

	if s.py > h {
		s.py = 0
	} else if s.py < 0 {
		s.py = h
	}

	s.head.Draw(screen)
	s.drawBody(screen)

	if s.stop {
		return
	}

	update := time.Now()
	delta := update.Sub(s.update).Seconds()

	if (delta <= 1) {
		return
	} else {
		s.update = update
	}

	x, y := s.head.Position()
	switch s.dir {
	case KeyArrowRight:
		x += 1
		break
	case KeyArrowDown:
		y += 1
		break
	case KeyArrowLeft:
		x -= 1
		break
	case KeyArrowUp:
		y -= 1
		break
	}

	s.head.SetPosition(x, y)
	s.px = x
	s.py = y
	s.moveBody()
}

func (s *Snake) Tick(event tl.Event) {
	x, y := s.px, s.py
	var d Direct
	var update bool

	if event.Type == tl.EventKey {
		switch event.Key {
			case tl.KeyArrowRight:
				x += 1
				d = KeyArrowRight
				update = true
				break 
			case tl.KeyArrowLeft:
				x -= 1
				d = KeyArrowLeft
				update = true
				break
			case tl.KeyArrowUp:
				y -= 1
				d = KeyArrowUp
				update = true
				break
			case tl.KeyArrowDown:
				y += 1
				d = KeyArrowDown
				update = true
				break
			case tl.KeySpace:
				if s.stop == true {
					s.stop = false
				} else {
					s.stop = true
				}
				s.update = time.Now()
				return
			default:
				d = s.dir
				break
		}

		if s.dir + d == 0 {
			return
		}
	} 

	if update {
		s.update = time.Now()
	}

	s.px = x 
	s.py = y
	s.dir = d
	s.head.SetPosition(x, y)
	s.moveBody()
}

func (s *Snake) Level () (*tl.BaseLevel) {
	return s.level
}

func (s *Snake) moveBody() {
	if s.size == 0 {
		return
	}

	d := s.body[0].dir
	s.body[0].Move(s.dir)
	for i := 1; i < s.size; i++ {
		d = s.body[i].Move(d)
	}
}

func (s *Snake) drawBody(screen *tl.Screen) {
	for i := 0; i < s.size; i++ {
		s.body[i].Draw(screen)
	}
}
