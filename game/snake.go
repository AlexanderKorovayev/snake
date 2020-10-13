package game

import (
	"github.com/JoelOtter/termloop"
)

//CreateSnake создать змейку
func CreateSnake() *Snake {
	snakeObj := new(Snake)
	snakeObj.Entity = termloop.NewEntity(1, 1, 1, 1)
	snakeObj.Direction = right
	snakeObj.Body = []coordinates{{1, 1}, {1, 2}, {1, 3}}
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

	//отрисовка на экране
	for _, v := range snake.Body {
		screen.RenderCell(v.x, v.y, &termloop.Cell{Fg: termloop.ColorWhite,
			Bg: termloop.ColorBlack,
			Ch: rune('߷')})
	}

}

// Tick позволяет сущности человека двигаться
// каждый момент времени мы проверяем состояние, что бы отрисовать объекты
func (snake *Snake) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		if event.Key == termloop.KeyArrowRight {
			snake.Direction = right
		}
		if event.Key == termloop.KeyArrowLeft {
			snake.Direction = left
		}
		if event.Key == termloop.KeyArrowUp {
			snake.Direction = up
		}
		if event.Key == termloop.KeyArrowDown {
			snake.Direction = down
		}
	}
}
