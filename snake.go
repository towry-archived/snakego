package main

import tl "github.com/JoelOtter/termloop"
// import "fmt"

type Direct int 

const BodyLength = 5

const (
	KeyArrowUp Direct = -1
	KeyArrowRight  Direct = -2
	KeyArrowLeft  Direct = 2
	KeyArrowDown  Direct = 1
)

type Snake struct {
	head *tl.Entity
	body [BodyLength]*Body
	food *Food
	level *tl.BaseLevel
	px int 
	py int
	dir Direct
	size int
}

func NewSnake(game *tl.Game) (*Snake) {	

	snake := new(Snake)
	snake.head = tl.NewEntity(1, 1, 1, 1)
	snake.px = BodyLength
	snake.py = 0
	snake.head.SetPosition(snake.px, snake.py)
	snake.head.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '#'})
	snake.dir = KeyArrowRight

	snake.level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorCyan,
	})

	snake.level.AddEntity(snake)

	// lets start with simple
	// suppose the snake is start from the top left corner
	for i := 0; i < BodyLength; i++ {
		snake.body[i] = NewBody(BodyLength - i - 1, 0, KeyArrowRight)
	}

	// create food
	snake.food = NewFood()
	snake.level.AddEntity(snake.food)

	return snake
}


func (s *Snake) Draw(screen *tl.Screen) {
	s.head.Draw(screen)
	s.drawBody(screen)
}

func (s *Snake) Size() (int, int) {
	return s.head.Size()
}

func (s *Snake) Position() (int, int) {
	return s.head.Position()
}

func (s *Snake) Collide(collision tl.Physical) {
	if _, ok := collision.(*Food); ok {
		s.food.entity.SetPosition(4, 4)
	}
}

func (s *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		x, y := s.px, s.py
		var d Direct

		switch event.Key {
			case tl.KeyArrowRight:
				x += 1
				d = KeyArrowRight
				break 
			case tl.KeyArrowLeft:
				x -= 1
				d = KeyArrowLeft
				break
			case tl.KeyArrowUp:
				y -= 1
				d = KeyArrowUp
				break
			case tl.KeyArrowDown:
				y += 1
				d = KeyArrowDown
				break
			default: 
				return
		}

		if s.dir + d == 0 {
			return
		}

		s.px = x 
		s.py = y
		s.dir = d
		s.head.SetPosition(x, y)

		s.moveBody()
	}
}

func (s *Snake) Level () (*tl.BaseLevel) {
	return s.level
}

func (s *Snake) moveBody() {
	d := s.body[0].dir
	s.body[0].Move(s.dir)
	for i := 1; i < BodyLength; i++ {
		d = s.body[i].Move(d)
	}
}

func (s *Snake) drawBody(screen *tl.Screen) {
	for i := 0; i < BodyLength; i++ {
		s.body[i].Draw(screen)
	}
}
