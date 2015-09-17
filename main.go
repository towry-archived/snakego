package main 

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()
	
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorCyan,
	})

	snake := NewSnake()

	// snake.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '□'})
	level.AddEntity(snake)

	game.Screen().SetLevel(level)
	game.Start()
}
