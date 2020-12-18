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
// он ведёт счётчик, который что бы считать клиентов
// инкриментируется при каждом запросе.
// Так же надо каким-то способом параллельно надо запустить обратный отсчёт

// как организовать обратный отсчёт
// создать два потока, один принимает запросы, а второй ведёт обратный отсчёт
// когда на первый приходит запрос он сначала идёт на второй что бы
// проверить есть ли время
// если времени больше нет, то при след запросах от клиента он
// посылает им собщение что время вышло и пора начинать

// InitiateGame ожидание игроков
func InitiateGame(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//добавить обработку ошибок
	}

	fmt.Println(string(body))

	var data core.TransportData
	err = json.Unmarshal(body, &data)

	if err != nil {
		//добавить обработку ошибок
	}

	// проверим, есть ли клиент в игре
	fmt.Println(reflect.TypeOf(data.Info))
	_, ok := core.ClientsCount[data.Info]
	// если клиент уже в игре
	if ok {
		//то посылаем ему информацию что идёт ожидание
	} else {
		// если лимит игроков не привышен
		if len(core.ClientsCount) < core.MaxObjectsCount {
			//добавляем клиента в список и посылаем ему информацию что идёт ожидание
		} else {
			//все игроки в сборе можно начинать досрочно
		}

	}

	myJSON, err := json.Marshal(data)
	if err != nil {
		//добавить обработку ошибок
	}
	fmt.Println(string(myJSON))
	//отправляем данные клиенту обратно
	fmt.Fprintf(w, string(myJSON))
}
