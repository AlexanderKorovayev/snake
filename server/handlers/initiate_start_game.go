/*
package handlers
модуль initiate_snake_coord
обработчик для определения начальных координат змейки
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

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
			fmt.Println(reflect.TypeOf(data.Info))
			_, ok := core.ClientsCount[data.Info]
			// если клиент уже в игре
			if ok {
				myJSON := addInfo(&data, "already added")
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			} else {
				// иначе добавляем в игру
				core.ClientsCount[data.Info] = nil
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
		// сообщаем клиенту, что время вышло
		data := parseBody(r)
		// если клиент уже был в игре, то отправим координаты всех объектов
		_, ok := core.ClientsCount[data.Info]
		// если клиент уже был в игре, то отправим координаты всех объектов
		if ok {
			//реализовать отправку всех объектов
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
	data.Info = string(core.TimeCount)
	// преобразуем данные в бинарный вид
	myJSON, err := json.Marshal(data)
	if err != nil {
		//добавить обработку ошибок
	}
	fmt.Println(string(myJSON))
	return myJSON
}

func parseBody(r *http.Request) core.TransportData {
	// если все условия соблюдены, то начинаем читать данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//добавить обработку ошибок
	}

	fmt.Println(string(body))
	// приводим данные к нужном формату
	var data core.TransportData
	err = json.Unmarshal(body, &data)

	if err != nil {
		//добавить обработку ошибок
	}
	return data
}
