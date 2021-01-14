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

// Direction направление змейки
type Direction int

const (
	// Up вверх
	Up Direction = iota
	// Down вниз
	Down
	// Left влево
	Left
	// Right вправо
	Right
)

// какие события могут быть?
// 1) добавить игрока - в ответ получаем координаты его змейки и доп инфу если есть
// 2) получить координаты объектов - в ответ получаем координаты всех объектов и доп инфу если есть
// TransportData структура, которая будет передаваться между сервером и клиентом
type TransportData struct {
	Action           interface{}              // create, getCoordinate. Подумать над статусами.
	MainObjectsCoord map[string][]Coordinates // тут будут координаты всех объектов
	Info             interface{}              // отсчёт для начал игры
	Color            string                   //цвет змейки
}
