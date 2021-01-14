/*
package core
модуль snake
описывает объект змейка
*/

package game

import (
	"encoding/json"

	"github.com/JoelOtter/termloop"
)

//CreateSnake создать змейку
func CreateSnake(body []Coordinates, name string, drctn direction, color termloop.Attr) *snake {
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
	if snake.areaCollision() || GameScreen.Snake1.dead == true {
		level := startFinishLevel()
		TermloopGame.Screen().SetLevel(level)
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
	message := new(TransportData)
	// зададим координаты змейки
	message.MainObjectsCoord = map[string][]Coordinates{GameScreen.Snake1.name: GameScreen.Snake1.body}
	// зададим направление змейки
	message.Action = GameScreen.Snake1.drctn
	// зададим имя змейки
	message.Info = GameScreen.Snake1.name
	// опрашиваем сервер
	info := getServerInfo("playersTurn", message)
	// распарсим info в json
	infoJSON := new(TransportData)
	err := json.Unmarshal(info, infoJSON)
	if err != nil {
		//добавить обработку ошибок
	}
	// если произошло столкновение, то остановим игру у клиента
	if infoJSON.Action.(string) == "snakeSelfCollision" {
		GameScreen.Snake1.dead = true
	}
	// обновим координаты для всех объектов
	for objName, coord := range infoJSON.MainObjectsCoord {
		if objName == "food" && GameScreen.GameFood != nil {
			GameScreen.GameFood.coord = coord[0]
		}
		if GameScreen.Snake1 != nil && objName == GameScreen.Snake1.name {
			GameScreen.Snake1.body = coord
		}
		if GameScreen.Snake2 != nil && objName == GameScreen.Snake2.name {
			GameScreen.Snake2.body = coord
		}
		if GameScreen.Snake3 != nil && objName == GameScreen.Snake3.name {
			GameScreen.Snake3.body = coord
		}
		if GameScreen.Snake4 != nil && objName == GameScreen.Snake4.name {
			GameScreen.Snake4.body = coord
		}
	}
}

//areaCollision определение коллизии с окружением
func (snake *snake) areaCollision() bool {
	return GameScreen.GameArea.collision(snake.GetHead())
}

//GetHead получение головы змейки
func (snake *snake) GetHead() *Coordinates {
	return &snake.body[len(snake.body)-1]
}
