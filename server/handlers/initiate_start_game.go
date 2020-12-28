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
				myJSON := addInfo(&data, "already added")
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			} else {
				// иначе добавляем в игру
				core.ClientsCount[data.Info.(string)] = ""
				myJSON := addInfo(&data, "added")
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			}
		} else {
			// сообщаем клиенту, что мест больше нет
			data := core.TransportData{}
			myJSON := addInfo(&data, "busy")
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
			// надо перебрать всех подключённых клиентов через ClientsCount
			for clName := range core.ClientsCount {
				//получаем координаты для данного клиента
				data.MainObjectsCoord[clName] = generateBodyCoord(i)
				i++
			}
			// зададим координаты для еды
			x, y := core.GetCoordinates()
			data.MainObjectsCoord["food"] = []core.Coordinates{{X: x, Y: y}}
			// сообщаем, что можно начинать играть
			// тестовые змейки
			data.MainObjectsCoord["192.168.1.144"] = []core.Coordinates{{X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}}
			data.MainObjectsCoord["192.168.1.145"] = []core.Coordinates{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}}
			myJSON := addInfo(&data, "ready")
			//отправляем данные клиенту обратно
			fmt.Printf("в итоге %v \n", data)
			fmt.Fprintf(w, string(myJSON))
		} else {
			// иначе сообщаем, что время для добавления вышло
			myJSON := addInfo(&data, "finished")
			//отправляем данные клиенту обратно
			fmt.Fprintf(w, string(myJSON))
		}
	}
}

func addInfo(data *core.TransportData, status string) []byte {
	// то посылаем ему информацию что идёт ожидание
	data.Action = status
	data.Info = strconv.Itoa(core.TimeCount)
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
func generateBodyCoord(numPlayer int) []core.Coordinates {
	var coord []core.Coordinates
	switch numPlayer {
	case 1:
		coord = []core.Coordinates{{1, core.High - 2}, {2, core.High - 2}, {3, core.High - 2}}
	case 2:
		coord = []core.Coordinates{{core.Width - 5, 2}, {core.Width - 4, 2}, {core.Width - 3, 2}}
	case 3:
		//допилить этот случай
		coord = []core.Coordinates{{1, core.High - 13}, {2, core.High - 13}, {3, core.High - 13}}
	case 4:
		//допилить этот случай
		coord = []core.Coordinates{{core.Width - 5, core.High - 2},
			{core.Width - 4, core.High - 2},
			{core.Width - 3, core.High - 2}}
	}
	return coord
}
