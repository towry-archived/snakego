package main

import tl "github.com/JoelOtter/termloop"

type Food struct {
	entity *tl.Entity
	x int
	y int
}

func NewFood() (*Food) {
	food := &Food{
		entity: tl.NewEntity(1, 1, 1, 1),
		x: -1,
		y: -1,
	}

	food.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '❁'})
	return food
}

func (food *Food) Draw(screen *tl.Screen) {
	if (food.x == -1) {
		sx, sw := screen.Size()
		rx := random(1, sx)
		ry := random(1, sw)
		food.x = rx
		food.y = ry
		food.entity.SetPosition(rx, ry)
	}

	food.entity.Draw(screen)
}

func (food *Food) Invalid() {
	food.x = -1
}

func (food *Food) Tick(event tl.Event) {
	// do not update this position
}

func (food *Food) Size() (int, int) {
	return food.entity.Size()
}

func (food *Food) Position() (int, int) {
	return food.entity.Position()
}

