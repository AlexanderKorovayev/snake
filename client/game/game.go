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
func startBaseSnakeLevel() *game {
	//создаём основные объекты
	gameScreen = new(game)
	gameScreen.Level = termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})
	// добавляем игровое поле
	gameScreen.gameArea = CreateArea()
	gameScreen.AddEntity(gameScreen.gameArea)

	return gameScreen
}

//startMainSnakeLevel формируем главную часть основного уровеня
func startMainSnakeLevel(objectsInfo map[string][]Coordinates, directionMap map[string]direction, color map[string]string) {
	snakesName := []string{"Snake2", "Snake3", "Snake4"}
	for objName, coord := range objectsInfo {
		if objName == "food" {
			// добавляем еду
			gameScreen.gameFood = CreateFood(coord[0])
			gameScreen.AddEntity(gameScreen.gameFood)
		} else {
			// у каждого клиента будет своя основная змейка,
			// у которой будут отслеживаться действия по тику.
			// Змейки других игроков будут отрисовываться как простые объекты.

			// добавляем змеек
			// главную змейку клиента всегда размещаем в Snake1
			if objName == clientID {
				gameScreen.snake1 = CreateSnake(coord, objName, directionMap[objName], colorMap(color[objName]))
				gameScreen.AddEntity(gameScreen.snake1)
			} else {
				// остальных змеек раскидываем по оставшимся местам
				// и просто отрисовываем
				for _, snakeName := range snakesName {
					if snakeName == "Snake2" {
						gameScreen.snake2 = CreateOtherSnake(coord, objName, colorMap(color[objName]))
						gameScreen.AddEntity(gameScreen.snake2)
						snakesName = remove(snakesName, "Snake2")
						break
					}
					if snakeName == "Snake3" {
						gameScreen.snake3 = CreateOtherSnake(coord, objName, colorMap(color[objName]))
						gameScreen.AddEntity(gameScreen.snake3)
						snakesName = remove(snakesName, "Snake3")
						break
					}
					if snakeName == "Snake4" {
						gameScreen.snake4 = CreateOtherSnake(coord, objName, colorMap(color[objName]))
						gameScreen.AddEntity(gameScreen.snake4)
						snakesName = remove(snakesName, "Snake4")
						break
					}
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
	TermloopGame.Screen().SetFps(4)
	TermloopGame.Screen().SetLevel(level)
	TermloopGame.Start()
}
