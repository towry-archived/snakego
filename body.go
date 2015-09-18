package main 

import tl "github.com/JoelOtter/termloop"

type Body struct {
	entity *tl.Entity
	px int
	py int
	dir Direct
}

func NewBody(x, y int, d Direct) (*Body) {
	body := &Body{
		entity: tl.NewEntity(1, 1, 1, 1),
		px: x,
		py: y,
		dir: d,
	}
	body.entity.SetPosition(x,y)
	body.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '#'})

	return body
}

func (body *Body) Draw(s *tl.Screen) {
	body.entity.Draw(s)
}

func (body *Body) Size() (int, int) {
	return body.entity.Size()
}

func (body *Body) Position() (int, int) {
	return body.entity.Position()
}

func (body *Body) Move(d Direct) (dir Direct) {
	var x, y int 
	x = body.px
	y = body.py

	switch body.dir {
		case KeyArrowRight:
			x += 1
			break 
		case KeyArrowUp:
			y -= 1
			break
		case KeyArrowLeft:
			x -= 1
			break 
		case KeyArrowDown:
			y += 1
			break
	}

	body.px = x 
	body.py = y 
	body.entity.SetPosition(x, y)
	dir = body.dir
	body.dir = d

	return dir
}
