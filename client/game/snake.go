/*
package core
модуль snake
описывает объект змейка
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//CreateSnake создать змейку
func CreateSnake(body []Coordinates) *snake {
	snakeObj := new(snake)
	snakeObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	snakeObj.drctn = right
	snakeObj.body = body
	return snakeObj
}

//Draw отвечает за отрисовку змеи на дисплее
func (snake *snake) Draw(screen *termloop.Screen) {
	if snake.drctn == right {
		head := snake.body[len(snake.body)-1]
		head.X++
		snake.body = append(snake.body[1:], head)
	}
	if snake.drctn == left {
		head := snake.body[len(snake.body)-1]
		head.X--
		snake.body = append(snake.body[1:], head)
	}
	if snake.drctn == up {
		head := snake.body[len(snake.body)-1]
		head.Y--
		snake.body = append(snake.body[1:], head)
	}
	if snake.drctn == down {
		head := snake.body[len(snake.body)-1]
		head.Y++
		snake.body = append(snake.body[1:], head)
	}

	if snake.foodCollision() {
		//увеличиваем длинну змейки
		snake.increaseSnake()
		//перемещаем еду на новое место
		GameScreen.GameFood.moveFood()
	}

	if snake.areaCollision() || snake.snakeCollision() {
		level := startFinishLevel()
		TermloopGame.Screen().SetLevel(level)
	}

	//отрисовка на экране
	for _, v := range snake.body {
		screen.RenderCell(v.X, v.Y, &termloop.Cell{Fg: termloop.ColorWhite,
			Bg: termloop.ColorWhite})
	}

}

// Tick позволяет отслеживать нажатия клавиатуры
func (snake *snake) Tick(event termloop.Event) {
	// теперь тик змейки у клиента становится основой работы с сервером.
	// Серверу будет отправляться нажатая клавиша и текущие координаты змейки.
	// Сервер в ответ посылает новые координаты змейки, координаты остальных объектов,
	// а так же событие столкновения с чем-либо, что бы клиент мог завершить игру.
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
