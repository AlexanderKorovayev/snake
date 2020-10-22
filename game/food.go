package game

import (
	"math/rand"
	"time"

	"github.com/JoelOtter/termloop"
)

//CreateFood создать еду для змейки
func CreateFood() *Food {
	foodObj := new(Food)
	foodObj.Entity = termloop.NewEntity(1, 1, 1, 1)

	x, y := GetCoordinates()
	// Set the new position of the food.
	foodObj.coord = Coordinates{x, y}

	return foodObj
}

//GetCoordinates получение координат для пищи
func GetCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси X
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	return r1.Intn(34), r2.Intn(19)
}

//Draw отвечает за отрисовку пищи на дисплее
func (food *Food) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(food.coord.x, food.coord.y, &termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
		Ch: rune('€')})
}

//Collision произошло ли косание с едой
func (food *Food) Collision(c *Coordinates) bool {
	return food.coord.x == c.x && food.coord.y == c.y
}

//MoveFood передвинуть еду
func (food *Food) MoveFood() {
	x, y := GetCoordinates()
	// Set the new position of the food.
	food.coord = Coordinates{x, y}
	//food.SetPosition(x, y)
}
