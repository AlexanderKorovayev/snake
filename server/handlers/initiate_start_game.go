/*
package handlers
модуль initiate_snake_coord
обработчик для инициализации игры
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		if len(core.ClientsCount) < core.MaxObjectsCount {
			// парсим входящие данные
			data := parseBody(r)
			// проверим, есть ли клиент в игре
			// с клиента будут приходить только строчки,
			// поэтому сделаем приведение типа
			_, ok := core.ClientsCount[data.Info.(string)]
			// если клиент уже в игре
			if ok {
				myJSON := addInfo(&data, "already added", false)
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			} else {
				// иначе добавляем в игру
				core.ClientsCount[data.Info.(string)] = ""
				myJSON := addInfo(&data, "added", false)
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			}
		} else {
			// сообщаем клиенту, что мест больше нет
			data := core.TransportData{}
			myJSON := addInfo(&data, "busy", false)
			//отправляем данные клиенту обратно
			fmt.Fprintf(w, string(myJSON))
		}
	} else {
		fmt.Println("время вышло")
		// сообщаем клиенту, что время вышло
		data := parseBody(r)
		_, ok := core.ClientsCount[data.Info.(string)]
		// если клиент уже был в игре, то отправим координаты всех объектов
		if ok {
			// отправка всех объектов
			fmt.Println("отправляем координаты")
			// введём порядковый номер, который нужен для
			// правильного распределения змеек
			i := 1
			// для каждого клиента будет создаваться заново
			core.DirectionMap = map[string]core.Direction{}
			// надо перебрать всех подключённых клиентов через ClientsCount
			for clName := range core.ClientsCount {
				//получаем координаты и направление для данного клиента
				data.MainObjectsCoord[clName], core.DirectionMap[clName] = generateDrctnBodyCoord(i)
				i++
			}
			// зададим координаты для еды
			x, y := core.GetCoordinates()
			data.MainObjectsCoord["food"] = []core.Coordinates{{X: x, Y: y}}
			// так же необходимо записать координаты и для сервера
			core.MainObjects = map[string][]core.Coordinates{} // сначала надо проинициализировать поле, но это можно сделать проще
			core.MainObjects["food"] = []core.Coordinates{{X: x, Y: y}}
			// сообщаем, что можно начинать играть
			// !!!!!!!!!!!!!!!!!!!тестовые змейки!!!!!!!!!!!!!!!!!!!
			//data.MainObjectsCoord["192.168.1.144"] = generateDrctnBodyCoord(2)
			//data.MainObjectsCoord["192.168.1.145"] = generateDrctnBodyCoord(3)
			//data.MainObjectsCoord["192.168.1.146"] = generateDrctnBodyCoord(4)
			myJSON := addInfo(&data, "ready", true)
			//отправляем данные клиенту обратно
			fmt.Printf("в итоге %v \n", data)
			fmt.Fprintf(w, string(myJSON))
		} else {
			// иначе сообщаем, что время для добавления вышло
			myJSON := addInfo(&data, "finished", false)
			//отправляем данные клиенту обратно
			fmt.Fprintf(w, string(myJSON))
		}
	}
}

func addInfo(data *core.TransportData, status string, addDrctn bool) []byte {
	// то посылаем ему информацию что идёт ожидание
	data.Action = status
	if addDrctn == true {
		data.Info = core.DirectionMap
	} else {
		data.Info = strconv.Itoa(core.TimeCount)
	}
	// преобразуем данные в бинарный вид
	myJSON, err := json.Marshal(data)
	if err != nil {
		//добавить обработку ошибок
	}
	return myJSON
}

func parseBody(r *http.Request) core.TransportData {
	// если все условия соблюдены, то начинаем читать данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//добавить обработку ошибок
	}

	// приводим данные к нужном формату
	var data core.TransportData
	err = json.Unmarshal(body, &data)

	if err != nil {
		//добавить обработку ошибок
	}
	return data
}

//generateBodyCoord генерируем координаты змейки для каждого игрока
//ставим каждого игрока в свой угол
func generateDrctnBodyCoord(numPlayer int) ([]core.Coordinates, core.Direction) {
	var coord []core.Coordinates
	var drctn core.Direction
	switch numPlayer {
	case 1:
		coord, drctn = []core.Coordinates{{X: 1, Y: core.High - 1}, {X: 2, Y: core.High - 1}, {X: 3, Y: core.High - 1}}, core.Right
	case 2:
		coord, drctn = []core.Coordinates{{X: core.Width - 3, Y: 1}, {X: core.Width - 2, Y: 1}, {X: core.Width - 1, Y: 1}}, core.Down
	case 3:
		coord, drctn = []core.Coordinates{{X: 1, Y: core.High - 14}, {X: 2, Y: core.High - 14}, {X: 3, Y: core.High - 14}}, core.Right
	case 4:
		coord, drctn = []core.Coordinates{{X: core.Width - 3, Y: core.High - 1},
			{X: core.Width - 2, Y: core.High - 1},
			{X: core.Width - 1, Y: core.High - 1}}, core.Left
	}
	return coord, drctn
}
