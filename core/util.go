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
func FindInSlice(sliceData *[]coordinates, data *coordinates) bool {
	for _, el := range *sliceData {
		if el == *data {
			return true
		}
	}
	return false
}

//GenerateBodyCoord генерируем координаты расположения змейки для каждого игрока
//распологаем каждого игрока в свой угол
func GenerateBodyCoord(numPlayer int) ([][]coordinates, error) {
	if numPlayer > maxPlayer {
		return nil, fmt.Errorf("большое колличество игроков, должно быть не более %v", numPlayer)
	} else {
		var bodys [][]coordinates
		switch numPlayer {
		case 1:
			bodys = [][]coordinates{{{33, 15}, {34, 15}, {35, 15}}}
		case 2:
			bodys = nil //доработать случай
		case 3:
			bodys = nil //доработать случай
		}
		return bodys, nil
	}
}
