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

//startBaseSnakeLevel формируем базовую часть основного уровеня
func startBaseSnakeLevel() *Game {
	//создаём основные объекты
	GameScreen = new(Game)
	GameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем игровое поле
	GameScreen.GameArea = CreateArea()
	GameScreen.AddEntity(GameScreen.GameArea)

	return GameScreen
}

//startMainSnakeLevel формируем главную часть основного уровеня
func startMainSnakeLevel(objectsInfo map[string][]Coordinates) {
	snakesName := []string{"Snake1", "Snake2", "Snake3", "Snake4"}
	for objName, coord := range objectsInfo {
		if objName == "food" {
			// добавляем еду
			logToFIle(coord)
			GameScreen.GameFood = CreateFood(coord[0])
			GameScreen.AddEntity(GameScreen.GameFood)
		} else {
			// добавляем змеек
			for _, snakeName := range snakesName {
				if snakeName == "Snake1" {
					GameScreen.Snake1 = CreateSnake(coord)
					GameScreen.AddEntity(GameScreen.Snake1)
					snakesName = remove(snakesName, "Snake1")
				}
				if snakeName == "Snake2" {
					GameScreen.Snake2 = CreateSnake(coord)
					GameScreen.AddEntity(GameScreen.Snake1)
					snakesName = remove(snakesName, "Snake2")
				}
				if snakeName == "Snake3" {
					GameScreen.Snake3 = CreateSnake(coord)
					GameScreen.AddEntity(GameScreen.Snake1)
					snakesName = remove(snakesName, "Snake3")
				}
				if snakeName == "Snake4" {
					GameScreen.Snake4 = CreateSnake(coord)
					GameScreen.AddEntity(GameScreen.Snake1)
					snakesName = remove(snakesName, "Snake4")
				}
			}
		}
	}
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
