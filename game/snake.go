package game

import (
	"github.com/JoelOtter/termloop"
)

//CreateSnake создать змейку
func CreateSnake() *Snake {
	snakeObj := new(Snake)
	snakeObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	snakeObj.Direction = right
	snakeObj.Body = []Coordinates{{1, 1}, {1, 2}, {1, 3}}
	return snakeObj
}

//Draw отвечает за отрисовку змеи на дисплее
func (snake *Snake) Draw(screen *termloop.Screen) {
	if snake.Direction == right {
		head := snake.Body[len(snake.Body)-1]
		head.x++
		snake.Body = append(snake.Body[1:], head)
	}
	if snake.Direction == left {
		head := snake.Body[len(snake.Body)-1]
		head.x--
		snake.Body = append(snake.Body[1:], head)
	}
	if snake.Direction == up {
		head := snake.Body[len(snake.Body)-1]
		head.y--
		snake.Body = append(snake.Body[1:], head)
	}
	if snake.Direction == down {
		head := snake.Body[len(snake.Body)-1]
		head.y++
		snake.Body = append(snake.Body[1:], head)
	}

	if snake.FoodCollision() {
		//увеличиваем длинну змейки
		snake.increaseSnake()
		//перемещаем еду на новое место
		GameScreen.gameFood.MoveFood()
	}

	if snake.AreaCollision() {
		GameOver()
	}

	//отрисовка на экране
	for _, v := range snake.Body {
		screen.RenderCell(v.x, v.y, &termloop.Cell{Fg: termloop.ColorWhite,
			Bg: termloop.ColorBlack,
			Ch: rune('߷')})
	}

}

// Tick позволяет сущности змейки двигаться
// каждый момент времени мы проверяем состояние, что бы отрисовать объекты
func (snake *Snake) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		if event.Key == termloop.KeyArrowRight {
			if snake.Direction != left {
				snake.Direction = right
			}
		}
		if event.Key == termloop.KeyArrowLeft {
			if snake.Direction != right {
				snake.Direction = left
			}
		}
		if event.Key == termloop.KeyArrowUp {
			if snake.Direction != down {
				snake.Direction = up
			}
		}
		if event.Key == termloop.KeyArrowDown {
			if snake.Direction != up {
				snake.Direction = down
			}
		}
	}
}

//GetHead получение головы змейки
func (snake *Snake) GetHead() *Coordinates {
	return &snake.Body[len(snake.Body)-1]
}

//FoodCollision определение коллизии с едой
func (snake *Snake) FoodCollision() bool {
	return GameScreen.gameFood.Collision(snake.GetHead())
}

//AreaCollision определение коллизии с окружением
func (snake *Snake) AreaCollision() bool {
	return GameScreen.gameArea.Collision(snake.GetHead())
}

func (snake *Snake) increaseSnake() {
	if snake.Direction == right {
		head := snake.Body[len(snake.Body)-1]
		head.x++
		snake.Body = append(snake.Body, head)
	}
	if snake.Direction == left {
		head := snake.Body[len(snake.Body)-1]
		head.x--
		snake.Body = append(snake.Body, head)
	}
	if snake.Direction == up {
		head := snake.Body[len(snake.Body)-1]
		head.y--
		snake.Body = append(snake.Body, head)
	}
	if snake.Direction == down {
		head := snake.Body[len(snake.Body)-1]
		head.y++
		snake.Body = append(snake.Body, head)
	}
}
