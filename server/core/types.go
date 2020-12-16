/*
package core
модуль types
содержит основные структуры данных необходимые для серверной части игры
*/

package core

//Coordinates координаты
type Coordinates struct {
	X int
	Y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

// какие события могут быть?
// 1) добавить игрока - в ответ получаем координаты его змейки и доп инфу если есть
// 2) получить координаты объектов - в ответ получаем координаты всех объектов и доп инфу если есть
// TransportData структура, которая будет передаваться между сервером и клиентом
type TransportData struct {
	Action           string                   // create, getCoordinate. Подумать над статусами.
	MainObjectsCoord map[string][]Coordinates // тут будут координаты всех объектов
	Info             string                   // отсчёт для начал игры
}
