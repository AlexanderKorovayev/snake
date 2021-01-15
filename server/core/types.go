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

// TransportData основная структура для передачи между клиентом и сервером
type TransportData struct {
	Estimate         string                   // обратный отсчёт для начала игры
	CLientDirection  Direction                // направление конкретной змейки
	DirectionMap     map[string]Direction     // направление змеек
	MainObjectsCoord map[string][]Coordinates // координаты всех объектов
	Info             string                   // полезная информация
	Color            map[string]string        // цвет змеек
	ClientID         string                   // идентификатор клиента
}
