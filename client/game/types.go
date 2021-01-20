/*
package core
модуль types
модуль содержит новые типы и их конструкторы, которые необходимы для удобства в работе
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

// transportData основная структура для передачи между клиентом и сервером
type transportData struct {
	Estimate         string                   // обратный отсчёт для начала игры
	CLientDirection  direction                // направление конкретной змейки
	DirectionMap     map[string]direction     // направление змеек
	MainObjectsCoord map[string][]Coordinates // координаты всех объектов
	Info             string                   // полезная информация
	Color            map[string]string        // цвет змеек
	ClientID         string                   // идентификатор клиента
}

// сreateTransportData конструктор основной структуры передачи данных
func createTransportData() *transportData {
	trData := new(transportData)
	trData.Estimate = ""
	trData.CLientDirection = direction(0)
	trData.DirectionMap = map[string]direction{}
	trData.MainObjectsCoord = map[string][]Coordinates{}
	trData.Info = ""
	trData.Color = map[string]string{}
	trData.ClientID = ""
	return trData
}
