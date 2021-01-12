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
func CreateSnake(body []Coordinates, name string) *snake {
	snakeObj := new(snake)
	snakeObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	snakeObj.drctn = right
	snakeObj.body = body
	snakeObj.name = name
	return snakeObj
}

//Draw отвечает за отрисовку змеи на дисплее
func (snake *snake) Draw(screen *termloop.Screen) {
	// это должно приходить с сервера
	/*
		if snake.foodCollision() {
			//увеличиваем длинну змейки
			snake.increaseSnake()
			//перемещаем еду на новое место
			GameScreen.GameFood.moveFood()
		}
		// это должно приходить с сервера
		if snake.areaCollision() || snake.snakeCollision() {
			level := startFinishLevel()
			TermloopGame.Screen().SetLevel(level)
		}
	*/
	//отрисовка на экране главной змейки клиента
	for _, v := range snake.body {
		screen.RenderCell(v.X, v.Y, &termloop.Cell{Fg: termloop.ColorWhite,
			Bg: termloop.ColorWhite})
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
	// обновим координаты для всех объектов
	for objName, coord := range infoJSON.MainObjectsCoord {
		if objName == "food" {
			GameScreen.GameFood.coord = coord[0]
		}
		if objName == GameScreen.Snake1.name {
			logToFIle("snake1")
			logToFIle(coord)
			GameScreen.Snake1.body = coord
		}
		if objName == GameScreen.Snake2.name {
			GameScreen.Snake2.body = coord
		}
		if objName == GameScreen.Snake3.name {
			GameScreen.Snake3.body = coord
		}
		if objName == GameScreen.Snake4.name {
			GameScreen.Snake4.body = coord
		}
	}
}

//GetHead получение головы змейки
func (snake *snake) GetHead() *Coordinates {
	return &snake.body[len(snake.body)-1]
}

//foodCollision определение коллизии с едой
func (snake *snake) foodCollision() bool {
	return GameScreen.GameFood.collision(snake.GetHead())
}

//areaCollision определение коллизии с окружением
func (snake *snake) areaCollision() bool {
	return GameScreen.GameArea.collision(snake.GetHead())
}

//snakeCollision определение столкновений змейки с самой собой
func (snake *snake) snakeCollision() bool {
	bodyWithoutHead := snake.body[:len(snake.body)-1]
	return FindInSlice(&bodyWithoutHead, snake.GetHead())
}

func (snake *snake) increaseSnake() {
	if snake.drctn == right {
		head := snake.body[len(snake.body)-1]
		head.X++
		snake.body = append(snake.body, head)
	}
	if snake.drctn == left {
		head := snake.body[len(snake.body)-1]
		head.X--
		snake.body = append(snake.body, head)
	}
	if snake.drctn == up {
		head := snake.body[len(snake.body)-1]
		head.Y--
		snake.body = append(snake.body, head)
	}
	if snake.drctn == down {
		head := snake.body[len(snake.body)-1]
		head.Y++
		snake.body = append(snake.body, head)
	}
}
