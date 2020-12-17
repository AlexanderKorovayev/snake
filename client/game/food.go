/*
package core
модуль food
содержит объекты, которые змейка может съесть
*/

package game

import (
	"math/rand"
	"time"

	"github.com/JoelOtter/termloop"
)

//CreateFood создать еду для змейки
func CreateFood() *food {
	food := new(food)
	food.Entity = termloop.NewEntity(1, 1, 1, 1)

	x, y := getCoordinates()
	//разместим еду на игровом поле
	food.coord = Coordinates{x, y}

	return food
}

//getCoordinates получение рандомных координат для пищи
func getCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси Y
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	x := r1.Intn(width - 1)
	if x == 0 {
		x++
	}
	y := r2.Intn(high - 1)
	if y == 0 {
		y++
	}
	return x, y
}

//Draw отвечает за отрисовку пищи на дисплее
func (food *food) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(food.coord.X,
		food.coord.Y,
		&termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
			Ch: rune('€')})
}

//collision произошло ли косание с едой
func (food *food) collision(c *Coordinates) bool {
	return food.coord.X == c.X && food.coord.Y == c.Y
}

//moveFood передвинуть еду
func (food *food) moveFood() {
	x, y := getCoordinates()
	//установить новые координаты для еды.
	food.coord = Coordinates{x, y}
}
