/*
package game
модуль game отвечает за инициализауию самой игры
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//startMenuLevel формируем стартовый уровень
func startMenuLevel() *startLevel {
	//создаём основные объекты
	gameScreen := new(startLevel)
	gameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем стартовую надпись
	gameScreen.startMenu = createStartMenu()
	gameScreen.AddEntity(gameScreen.startMenu)
	return gameScreen
}

//startFinishLevel формируем финишный уровень
func startFinishLevel() *finishLevel {
	//создаём основные объекты
	gameScreen := new(finishLevel)
	gameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем финишнюю надпись
	gameScreen.finishMenu = createFinishMenu()
	gameScreen.AddEntity(gameScreen.finishMenu)
	return gameScreen
}

//startSnakeLevel формируем основной уровень
func startSnakeLevel() *Game {
	//создаём основные объекты
	GameScreen = new(Game)
	GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем игровое поле
	GameScreen.GameArea = CreateArea()
	GameScreen.AddEntity(GameScreen.GameArea)

	// отправляем запрос серверу на получение координат для змейки
	body := getSnakeCoord()
	logToFIle(body)
	// продумать как мы будем получать несколько змеек
	//и создавать их на одном поле
	GameScreen.Snake = CreateSnake(body)
	GameScreen.AddEntity(GameScreen.Snake)

	// добавляем еду

	GameScreen.GameFood = CreateFood()
	GameScreen.AddEntity(GameScreen.GameFood)

	return GameScreen
}

//StartGame стартуем игру
func StartGame() {
	TermloopGame = termloop.NewGame()

	// создаём стартовый уровень
	level := startMenuLevel()
	TermloopGame.Screen().SetFps(5)
	TermloopGame.Screen().SetLevel(level)
	TermloopGame.Start()
}
