/*
package handlers
модуль initiate_snake_coord
обработчик для инициализации игры
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlexanderKorovaev/snake/server/core"
)

// каждый клиент делает запрос на этот обработчик
// он ведёт счётчик подключённых клиентов, который
// инкриментируется при каждом запросе.
// а так же проверяет, осталось ли время для подключений.
// и посылает коиенту инфу, что игра уже занята.

// InitiateGame ожидание игроков
func InitiateGame(w http.ResponseWriter, r *http.Request) {
	// проверим осталось ли время
	if core.TimeCount > 0 {
		// проверим остались ли свободные места
		if len(core.ColorMap) < core.MaxObjectsCount {
			// парсим входящие данные
			data := core.ParseBody(r)
			// проверим, есть ли клиент в игре
			_, ok := core.ColorMap[data.ClientID]
			// если клиент уже в игре
			if ok {
				myJSON := addInfo(data, "already added", core.ColorMap)
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			} else {
				// назначим цвет
				// так же обозначим что клиент добавлен в игру
				core.ColorMap[data.ClientID] = core.Colors[0]
				// удалим выбранный цвет из списка доступных
				core.Colors = core.Remove(core.Colors, core.Colors[0])
				myJSON := addInfo(data, "added", core.ColorMap)
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			}
		} else {
			// сообщаем клиенту, что мест больше нет
			data := core.CreateTransportData()
			myJSON := addInfo(data, "busy", core.ColorMap)
			//отправляем данные клиенту обратно
			fmt.Fprintf(w, string(myJSON))
		}
	} else {
		// сообщаем клиенту, что время вышло
		data := core.ParseBody(r)
		_, ok := core.ColorMap[data.ClientID]
		// если клиент уже был в игре, то отправим координаты всех объектов
		if ok {
			// отправка всех объектов
			// введём порядковый номер, который нужен для
			// правильного распределения змеек
			i := 1
			// надо перебрать всех подключённых клиентов через ColorMap
			for clName := range core.ColorMap {
				//получаем координаты и направление для данного клиента
				data.MainObjectsCoord[clName], data.DirectionMap[clName] = generateDrctnBodyCoord(i)
				i++
			}
			// зададим координаты для еды
			x, y := core.GetCoordinates()
			data.MainObjectsCoord["food"] = []core.Coordinates{{X: x, Y: y}}
			// запишем координаты еды на сервере для остальных клиентов
			core.MainObjects["food"] = []core.Coordinates{{X: x, Y: y}}
			// сообщаем, что можно начинать играть
			myJSON := addInfo(data, "ready", core.ColorMap)
			//отправляем данные клиенту обратно
			fmt.Fprintf(w, string(myJSON))
		} else {
			// иначе сообщаем, что время для добавления вышло
			myJSON := addInfo(data, "finished", core.ColorMap)
			//отправляем данные клиенту обратно
			fmt.Fprintf(w, string(myJSON))
		}
	}
}

func addInfo(data *core.TransportData, status string, colorMap map[string]string) []byte {
	// то посылаем ему информацию что идёт ожидание
	data.Info = status
	data.Estimate = strconv.Itoa(core.TimeCount)
	data.ColorMap = colorMap
	// преобразуем данные в бинарный вид
	myJSON, err := json.Marshal(data)
	if err != nil {
		//добавить обработку ошибок
	}
	return myJSON
}

//generateBodyCoord генерируем координаты змейки для каждого игрока
//ставим каждого игрока в свой угол
func generateDrctnBodyCoord(numPlayer int) ([]core.Coordinates, core.Direction) {
	var coord []core.Coordinates
	var drctn core.Direction
	switch numPlayer {
	case 1:
		coord, drctn = []core.Coordinates{{X: 1, Y: core.High - 1},
			{X: 2, Y: core.High - 1},
			{X: 3, Y: core.High - 1}}, core.Right
	case 2:
		coord, drctn = []core.Coordinates{{X: core.Width - 1, Y: 1},
			{X: core.Width - 2, Y: 1},
			{X: core.Width - 3, Y: 1}}, core.Left
	case 3:
		coord, drctn = []core.Coordinates{{X: 1, Y: core.High - 14},
			{X: 2, Y: core.High - 14},
			{X: 3, Y: core.High - 14}}, core.Right
	case 4:
		coord, drctn = []core.Coordinates{{X: core.Width - 1, Y: core.High - 1},
			{X: core.Width - 2, Y: core.High - 1},
			{X: core.Width - 3, Y: core.High - 1}}, core.Left
	}
	return coord, drctn
}
