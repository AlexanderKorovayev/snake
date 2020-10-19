package game

import (
	"github.com/JoelOtter/termloop"
)

//StartGame стартуем игру
func StartGame() {
	sg := termloop.NewGame()

	// Create titlescreen.
	gs := NewGamescreen()
	sg.Screen().SetFps(5)
	sg.Screen().SetLevel(gs)
	sg.Start()
}

//NewGamescreen инициализация игры
func NewGamescreen() *Game {
	// Creates the gamescreen level and create the entities
	gs := new(Game)
	gs.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	gs.snake = CreateSnake()
	// Add entities for the game level.
	gs.AddEntity(gs.snake)

	gs.gameArea = CreateArea()
	// Add entities for the game level.
	gs.AddEntity(gs.gameArea)

	return gs
}
