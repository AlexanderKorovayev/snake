/*
package core
модуль util
содержит вспомогательные функции
*/

package core

import (
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
