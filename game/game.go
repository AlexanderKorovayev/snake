/*
package game
модуль game отвечает за инициализауию самой игры
*/

package game

import (
	"github.com/AlexanderKorovaev/snake/core"
	"github.com/JoelOtter/termloop"
)

//startLevel инициализация уровня
func startLevel() *StartLevel {
	//создаём основные объекты
	GameScreen := new(StartLevel)
	GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем стартовую надпись
	GameScreen.StartMenu = CreateStartMenu()
	GameScreen.AddEntity(GameScreen.StartMenu)
	return GameScreen
}

//snakeLevel инициализация уровня
func snakeLevel() *core.Game {
	//создаём основные объекты
	core.GameScreen = new(core.Game)
	core.GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем игровое поле
	core.GameScreen.GameArea = core.CreateArea()
	core.GameScreen.AddEntity(core.GameScreen.GameArea)
	// добавляем змеек, пропорционально количетсву игроков
	players := 1 //подумать где и как будем брать игроков
	bodys, err := core.GenerateBodysCoord(players)
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
	core.TermloopGame = termloop.NewGame()

	// создаём стартовый уровень
	level := startLevel()
	core.TermloopGame.Screen().SetFps(5)
	core.TermloopGame.Screen().SetLevel(level)
	core.TermloopGame.Start()
}
