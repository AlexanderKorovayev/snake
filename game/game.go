/*
package game
модуль game отвечает за инициализауию самой игры
*/

package game

import (
	"github.com/AlexanderKorovaev/snake/core"
	"github.com/JoelOtter/termloop"
)

//NewLevel инициализация уровня
func NewLevel() *core.Game {
	//создаём основные объекты
	core.GameScreen = new(core.Game)
	core.GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем игровое поле
	core.GameScreen.GameArea = core.CreateArea()
	core.GameScreen.AddEntity(core.GameScreen.GameArea)
	// добавляем змеек, соответствующе количетсву игроков
	players := 1 //подумать где и как будем брать игроков
	bodys, err := core.GenerateBodyCoord(players)
	if err != nil {
		panic(err.Error())
	}
	for _, body := range bodys {
		core.GameScreen.Snake = core.CreateSnake(body)
		core.GameScreen.AddEntity(core.GameScreen.Snake)
	}
	// добавляем еду
	core.GameScreen.GameFood = core.CreateFood()
	core.GameScreen.AddEntity(core.GameScreen.GameFood)

	return core.GameScreen
}

//StartGame стартуем игру
func StartGame() {
	game := termloop.NewGame()

	// создаём основной экран
	level := NewLevel()
	game.Screen().SetFps(5)
	game.Screen().SetLevel(level)
	game.Start()
}
