/*
package core
модуль util
содержит вспомогательные функции
*/

package core

//FindInSlice функция для поиска вхождений в срезе
func FindInSlice(sliceData *[]Coordinates, data *Coordinates) bool {
	for _, el := range *sliceData {
		if el == *data {
			return true
		}
	}
	return false
}
