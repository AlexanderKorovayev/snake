/*
package game
модуль game отвечает за инициализауию самой игры
*/

package game

import (
	"encoding/json"

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
	// надо опрашивать сервер 10 сек с промежутком в секунду.
	// сначала змейка не двигается и на каждом опросе мы рисуем обратный отсчёт
	// силами клиента и когда видим событие финиш от сервера то змейки начинают ползти

	// отправляем серваку свою готовность играть.
	// внутри себя сервак запускает обратный отсчёт на добавление
	// остальных игроков и будет ждать только это время.
	// по истечению этого времени он отошлёт сообщение о готовности играть
	// а также координаты для всех объектов.
	// клиент в бесконечном цикле опрашивает сервер и если в ответе число
	// то отрисовываем его, если в ответе ready, то рисуем все объекты и
	// дальше по тику делаем запросы на перерисовку
	// всех объектов получаем координаты.

	// цикл опроса готовности сервера к игре
	for {
		// опрашиваем сервер
		info := getServerInfo()

		// распарсим info в json
		infoJSON := new(TransportData)
		infoJSON.MainObjectsCoord = map[string][]Coordinates{}
		err := json.Unmarshal(info, infoJSON)

		if err != nil {
			//добавить обработку ошибок
		}

		// теперь надо добавить проверку infoJSON на то что внутри
		// вытащим координаты из результата
		parseSnakeCoord(infoJSON)
	}

	// решить вопрос, как будем создавать змеек, переменная-то одна
	GameScreen.Snake = CreateSnake(info)
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
