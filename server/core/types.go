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
	ColorMap         map[string]string        // цвет змеек
	ClientID         string                   // идентификатор клиента
}

// CreateTransportData конструктор основной структуры передачи данных
func CreateTransportData() *TransportData {
	trData := new(TransportData)
	trData.Estimate = ""
	trData.CLientDirection = Direction(0)
	trData.DirectionMap = map[string]Direction{}
	trData.MainObjectsCoord = map[string][]Coordinates{}
	trData.Info = ""
	trData.ColorMap = map[string]string{}
	trData.ClientID = ""
	return trData
}
