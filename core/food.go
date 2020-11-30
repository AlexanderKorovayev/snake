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

	x, y := getCoordinates()
	//разместим еду на игровом поле
	food.coord = coordinates{x, y}

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
func (food *Food) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(food.coord.x,
		food.coord.y,
		&termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
			Ch: rune('€')})
}

//collision произошло ли косание с едой
func (food *Food) collision(c *coordinates) bool {
	return food.coord.x == c.x && food.coord.y == c.y
}

//moveFood передвинуть еду
func (food *Food) moveFood() {
	x, y := getCoordinates()
	//установить новые координаты для еды.
	food.coord = coordinates{x, y}
}
