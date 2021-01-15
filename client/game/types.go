/*
package core
модуль types
модуль содержит новые типы, которые необходимы для удобства в работе
*/

package game

// сделан доступным так как для приведения к json нужно
// что бы объекты были открытыми
// Сoordinates координаты
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

// TransportData основная структура для передачи между клиентом и сервером
type TransportData struct {
	Estimate         string                   // обратный отсчёт для начала игры
	CLientDirection  direction                // направление конкретной змейки
	DirectionMap     map[string]direction     // направление змеек
	MainObjectsCoord map[string][]Coordinates // координаты всех объектов
	Info             string                   // полезная информация
	Color            map[string]string        // цвет змеек
	ClientID         string                   // идентификатор клиента
}
