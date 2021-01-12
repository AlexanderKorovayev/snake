/*
package handlers
модуль players_turn
обработчик для просчёта действий всех игроков за один ход
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AlexanderKorovaev/snake/server/core"
)

// PlayersTurn получить координаты всех объектов
func PlayersTurn(w http.ResponseWriter, r *http.Request) {
	// клиент присылает координаты своей змейки и её направление
	// надо организовать обработку этой информации.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//добавить обработку ошибок
	}
	var data core.TransportData
	err = json.Unmarshal(body, &data)

	if err != nil {
		//добавить обработку ошибок
	}

	// теперь необходимо записывать данные каждого клиента.
	// клиенты будут запрашивать обновление данных, каждому мы будем присылать
	// последние актуальные данные.
	// так же необходимо каждый раз просчитвать, было ли столкновение с едой.

	// запишем координаты для поступившего запроса от клиента
	fmt.Println(data.MainObjectsCoord)
	clName := (data.Info).(string)
	core.MainObjects[clName] = data.MainObjectsCoord[clName]
	// просчитаем для данного клиента новые координаты
	// сначала приводим интерфейс к типу флоат а потом его уже к типу дирекшен
	drctn := core.Direction((data.Action).(float64))
	core.MainObjects[clName] = updateSnakeCoordinates(core.MainObjects[clName], drctn)
	// запишем клиенту все координаты
	data.MainObjectsCoord = core.MainObjects
	// отправляем данные клиенту
	myJSON, err := json.Marshal(data)
	if err != nil {
		//добавить обработку ошибок
	}
	//отправляем данные клиенту обратно
	fmt.Fprintf(w, string(myJSON))
}

// обновление координат змейки
func updateSnakeCoordinates(body []core.Coordinates, drctn core.Direction) []core.Coordinates {
	if drctn == core.Right {
		head := body[len(body)-1]
		head.X++
		body = append(body[1:], head)
	}
	if drctn == core.Left {
		head := body[len(body)-1]
		head.X--
		body = append(body[1:], head)
	}
	if drctn == core.Up {
		head := body[len(body)-1]
		head.Y--
		body = append(body[1:], head)
	}
	if drctn == core.Down {
		head := body[len(body)-1]
		head.Y++
		body = append(body[1:], head)
	}
	return body
}
