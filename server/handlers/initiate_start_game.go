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
			fmt.Printf("пришло %v \n", data)

			_, ok := core.ClientsCount[data.Info.(string)]
			fmt.Printf("есть ли такой игрок: %v \n", ok)
			// если клиент уже в игре
			if ok {
				fmt.Println("in")
				myJSON := addInfo(&data, "already added")
				fmt.Println(string(myJSON))
				//отправляем данные клиенту обратно
				fmt.Fprintf(w, string(myJSON))
			} else {
				fmt.Println("in1")
				// иначе добавляем в игру
				core.ClientsCount[data.Info.(string)] = ""
				myJSON := addInfo(&data, "added")
				fmt.Println(string(myJSON))
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
		_, ok := core.ClientsCount[data.Info.(string)]
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
	data.Info = core.TimeCount
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
