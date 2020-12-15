/*
package core
модуль util
содержит вспомогательные функции
*/

package core

import (
	"fmt"
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

//GenerateBodysCoord генерируем координаты змейки для каждого игрока
//ставим каждого игрока в свой угол
func GenerateBodysCoord(numPlayer int) ([][]Coordinates, error) {
	if numPlayer > maxPlayer {
		errTxt := "большое колличество игроков, должно быть не более %v"
		return nil, fmt.Errorf(errTxt, maxPlayer)
	} else {
		var bodys [][]Coordinates
		switch numPlayer {
		case 1:
			bodys = [][]Coordinates{{{1, high - 2}, {2, high - 2}, {3, high - 2}}}
		case 2:
			//доработать случай
		case 3:
			//доработать случай
		case 4:
			//доработать случай
		}
		return bodys, nil
	}
}
