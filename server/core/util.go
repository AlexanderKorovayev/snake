/*
package core
модуль util
содержит вспомогательные функции
*/

package core

import (
	"fmt"
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
		fmt.Println("in0")
		TimeCount--
		fmt.Println("in1")
		time.Sleep(time.Second)
		fmt.Println("in2")
	}
}
