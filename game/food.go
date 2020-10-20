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

	x, y := getCoordinates()
	// Set the new position of the food.
	foodObj.coord = coordinates{x, y}

	return foodObj
}

func getCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси X
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	return r1.Intn(35), r2.Intn(20)
}

//Draw отвечает за отрисовку змеи на дисплее
func (food *Food) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(food.coord.x, food.coord.y, &termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
		Ch: rune('€')})
}
