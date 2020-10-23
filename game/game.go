package game

import (
	"io/ioutil"

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
	GameScreen = new(Game)
	GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	GameScreen.snake = CreateSnake()
	// Add entities for the game level.
	GameScreen.AddEntity(GameScreen.snake)

	GameScreen.gameArea = CreateArea()
	// Add entities for the game level.
	GameScreen.AddEntity(GameScreen.gameArea)

	GameScreen.gameFood = CreateFood()
	// Add entities for the game level.
	GameScreen.AddEntity(GameScreen.gameFood)

	return GameScreen
}

//GameOver функция для отображения конца игры
func GameOver() {
	GameScreen.Level.RemoveEntity(GameScreen.snake)
	GameScreen.Level.RemoveEntity(GameScreen.gameArea)
	GameScreen.Level.RemoveEntity(GameScreen.gameFood)

	dat, _ := ioutil.ReadFile("gameover-logo.txt")
	e := termloop.NewEntityFromCanvas(1, 1, termloop.CanvasFromString(string(dat)))
	GameScreen.AddEntity(e)
}
