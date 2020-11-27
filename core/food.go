/*
package core
модуль food
содержит объекты, которые змейка может съесть
*/

package core

import (
	"math/rand"
	"time"

	"github.com/JoelOtter/termloop"
)

//CreateFood создать еду для змейки
func CreateFood() *Food {
	food := new(Food)
	food.Entity = termloop.NewEntity(1, 1, 1, 1)

	x, y := GetCoordinates()
	//разместим еду на игровом поле
	food.coord = Coordinates{x, y}

	return food
}

//GetCoordinates получение рандомных координат для пищи
func GetCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси Y
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	return r1.Intn(34), r2.Intn(19)
}

//Draw отвечает за отрисовку пищи на дисплее
func (food *Food) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(food.coord.X,
		food.coord.Y,
		&termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
			Ch: rune('€')})
}

//Collision произошло ли косание с едой
func (food *Food) Collision(c *Coordinates) bool {
	return food.coord.X == c.X && food.coord.Y == c.Y
}

//MoveFood передвинуть еду
func (food *Food) MoveFood() {
	x, y := GetCoordinates()
	//установить новые координаты для еды.
	food.coord = Coordinates{x, y}
}
