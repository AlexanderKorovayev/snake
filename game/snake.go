package game

import (
	tl "github.com/JoelOtter/termloop"
)

// переменная, которая хранит направление змейки
var snakeRout string = "Up"

// Snake сущность змейка
// должна наследовать сущность, что бы иметь возможность добавления на уровень
type Snake struct {
	*tl.Entity
	prevX int // нужен для откатывания при столкновениях
	prevY int // нужен для откатывания при столкновениях
	level *tl.BaseLevel
}

//Draw позволяет двигать камеру за персонажем
func (snake *Snake) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := snake.Position()
	snake.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	snake.Entity.Draw(screen)
}

// функция для определения движения змейки в тех случаях,
// когда никаких событий не происходит
func defaultTick(snake *Snake, event tl.Event) {
	snake.prevX, snake.prevY = snake.Position()
	if snakeRout == "Right" {
		snake.SetPosition(snake.prevX+1, snake.prevY)
	}
	if snakeRout == "Left" {
		snake.SetPosition(snake.prevX-1, snake.prevY)
	}
	if snakeRout == "Up" {
		snake.SetPosition(snake.prevX, snake.prevY-1)
	}
	if snakeRout == "Down" {
		snake.SetPosition(snake.prevX, snake.prevY+1)
	}
}

// Tick позволяет сущности человека двигаться
// каждый момент времени мы проверяем состояние, что бы отрисовать объекты
func (snake *Snake) Tick(event tl.Event) {
	if event.Type != tl.EventKey {
		// движение по умолчанию
		defaultTick(snake, event)
	} else {
		snake.prevX, snake.prevY = snake.Position()
		if event.Key == tl.KeyArrowRight {
			snake.SetPosition(snake.prevX+1, snake.prevY)
			snakeRout = "Right"
		}
		if event.Key == tl.KeyArrowLeft {
			snake.SetPosition(snake.prevX-1, snake.prevY)
			snakeRout = "Left"
		}
		if event.Key == tl.KeyArrowUp {
			snake.SetPosition(snake.prevX, snake.prevY-1)
			snakeRout = "Up"
		}
		if event.Key == tl.KeyArrowDown {
			snake.SetPosition(snake.prevX, snake.prevY+1)
			snakeRout = "Down"
		}
	}
}

// Collide позволяет сущности человека взаимодействовать с другими объектами
func (snake *Snake) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		snake.SetPosition(snake.prevX, snake.prevY)
	}
}
