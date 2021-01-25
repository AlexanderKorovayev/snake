/*
package core
модуль snake
описывает объект змейка
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//createSnake создать змейку
func createSnake(body []coordinates, name string, drctn direction, color termloop.Attr) *snake {
	snakeObj := new(snake)
	snakeObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	snakeObj.drctn = drctn
	snakeObj.body = body
	snakeObj.name = name
	snakeObj.dead = false
	snakeObj.color = color
	return snakeObj
}

//Draw отвечает за отрисовку змеи на дисплее
func (snake *snake) Draw(screen *termloop.Screen) {
	// столкновения с уровнем будут просчитываться на клиенте
	if snake.areaCollision() || gameScreen.snake1.dead == true {
		level := startFinishLevel()
		termloopGame.Screen().SetLevel(level)
	}
	//отрисовка на экране главной змейки клиента
	for _, v := range snake.body {
		screen.RenderCell(v.X, v.Y, &termloop.Cell{Fg: termloop.ColorWhite,
			Bg: snake.color})
	}

}

// Tick позволяет отслеживать нажатия клавиатуры
func (snake *snake) Tick(event termloop.Event) {
	// Теперь тик змейки у клиента становится основой работы с сервером.
	// Серверу будет отправляться нажатая клавиша и текущие координаты змейки.
	// Сервер в ответ посылает новые координаты змейки, координаты остальных объектов,
	// а так же событие столкновения с чем-либо, что бы клиент мог завершить игру.

	// Так же по тику мы будем обновлять snake.body, который сейчас в drow обовляется
	// и snake.body будет уже отрисовываться в drow как и раньше.

	// сначала пытаемся получить нажатия клавиш
	if event.Type == termloop.EventKey {
		if event.Key == termloop.KeyArrowRight {
			if snake.drctn != left {
				snake.drctn = right
			}
		}
		if event.Key == termloop.KeyArrowLeft {
			if snake.drctn != right {
				snake.drctn = left
			}
		}
		if event.Key == termloop.KeyArrowUp {
			if snake.drctn != down {
				snake.drctn = up
			}
		}
		if event.Key == termloop.KeyArrowDown {
			if snake.drctn != up {
				snake.drctn = down
			}
		}
	}

	// теперь получаем данные от сервера, что бы обработать актуальную
	// информацию о других объектах.

	// создадим сообщение, которое необходимо передать серверу
	message := createTransportData()
	// зададим координаты змейки
	// каждый клиент отправляет только свою главную змейку
	message.MainObjectsCoord = map[string][]coordinates{gameScreen.snake1.name: gameScreen.snake1.body}
	// зададим направление змейки
	message.CLientDirection = gameScreen.snake1.drctn
	// зададим имя змейки
	message.ClientID = gameScreen.snake1.name
	// опрашиваем сервер
	infoByte := getServerInfo("playersTurn", message)
	// распарсим info в json
	message = parseBody(infoByte)
	// если произошло столкновение, то остановим игру у клиента
	if message.Info == "snakeSelfCollision" {
		gameScreen.snake1.dead = true
	}
	// обновим координаты для всех объектов
	for objName, coord := range message.MainObjectsCoord {
		if objName == "food" && gameScreen.gameFood != nil {
			gameScreen.gameFood.coord = coord[0]
		}
		if gameScreen.snake1 != nil && objName == gameScreen.snake1.name {
			gameScreen.snake1.body = coord
		}
		if gameScreen.snake2 != nil && objName == gameScreen.snake2.name {
			gameScreen.snake2.body = coord
		}
		if gameScreen.snake3 != nil && objName == gameScreen.snake3.name {
			gameScreen.snake3.body = coord
		}
		if gameScreen.snake4 != nil && objName == gameScreen.snake4.name {
			gameScreen.snake4.body = coord
		}
	}
}

//areaCollision определение коллизии с окружением
func (snake *snake) areaCollision() bool {
	return gameScreen.gameArea.collision(snake.GetHead())
}

//GetHead получение головы змейки
func (snake *snake) GetHead() *coordinates {
	return &snake.body[len(snake.body)-1]
}
