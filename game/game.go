/*
package game
модуль game отвечает за инициализауию самой игры
*/

package game

import (
	"github.com/AlexanderKorovaev/snake/core"
	"github.com/JoelOtter/termloop"
)

//NewGamescreen инициализация игры
func NewGamescreen() *core.Game {
	//создаём основные объекты
	GameScreen = new(core.Game)
	GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})

	GameScreen.gameArea = core.CreateArea()
	// добавляем игровое поле
	GameScreen.AddEntity(GameScreen.gameArea)

	GameScreen.snake = core.CreateSnake()
	// добавляем змейку
	GameScreen.AddEntity(GameScreen.snake)

	GameScreen.gameFood = core.CreateFood()
	// добавляем еду
	GameScreen.AddEntity(GameScreen.gameFood)

	return GameScreen
}

//StartGame стартуем игру
func StartGame() {
	sg := termloop.NewGame()

	// создаём основной экран
	gs := NewGamescreen()
	sg.Screen().SetFps(5)
	sg.Screen().SetLevel(gs)
	sg.Start()
}
