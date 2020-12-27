/*
package core
модуль util
содержит вспомогательные функции
*/

package core

import (
	"math/rand"
	"time"
)

//FindInSlice функция для поиска вхождений в срезе
func FindInSlice(sliceData *[]Coordinates, data *Coordinates) bool {
	for _, el := range *sliceData {
		if el == *data {
			return true
		}
	}
	return false
}

// Countdown функция для обратного отсчёта перед началом игры
func Countdown() {
	for TimeCount > 0 {
		//fmt.Println(TimeCount)
		TimeCount--
		time.Sleep(time.Second)
	}
}

//GetCoordinates получение рандомных координат для пищи
func GetCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси Y
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	x := r1.Intn(Width - 1)
	if x == 0 {
		x++
	}
	y := r2.Intn(High - 1)
	if y == 0 {
		y++
	}
	return x, y
}
