package main 

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()
	
	snake := NewSnake(game)

	game.Screen().SetLevel(snake.Level())
	game.Start()
}
