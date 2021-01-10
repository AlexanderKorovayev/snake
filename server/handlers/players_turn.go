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

// GetSnakeCoordHandler получить координаты змейки
func GetSnakeCoordHandler(w http.ResponseWriter, r *http.Request) {
	// клиент присылает координаты своего объекта, надо организовать обработку дальнейших дейтсвий.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//добавить обработку ошибок
	}

	fmt.Println(string(body))

	//пробую создать необходимую структуру
	//testData := core.TransportData{"test", map[string][]core.Coordinates{"test": {{1, 1}}}, "test1"}
	//fmt.Println(testData)

	var data core.TransportData
	// можно делать универсальные подходы при парсинге json
	// var testResp interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		//добавить обработку ошибок
	}

	// тип интерфейс, но к нему нельзя обращаться по индексу, поэтому необходимо
	// привести явно к типу
	// тоже самое придётся сделать и со значениями по ключу, потому что они
	// тоже типа интерфейс
	//testRespTyped := testResp.(map[string]interface{})
	//fmt.Println(testRespTyped)

	//получаем координаты для поступившего запроса
	clName := (data.Info).(string)
	// прибавляем единицу, что бы внести больше понятности.
	// если длинна контейнера ноль, то значит это будет первый клиент
	data.MainObjectsCoord[clName] = generateBodyCoord(len(data.MainObjectsCoord) + 1)

	myJSON, err := json.Marshal(data)
	if err != nil {
		//добавить обработку ошибок
	}
	fmt.Println(string(myJSON))
	//отправляем данные клиенту обратно
	fmt.Fprintf(w, string(myJSON))
}

/*
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
		coord = []core.Coordinates{{core.Width - 5, 2}, {core.Width - 4, 2}, {core.Width - 3, 2}}
	case 4:
		//допилить этот случай
		coord = []core.Coordinates{{core.Width - 5, 2}, {core.Width - 4, 2}, {core.Width - 3, 2}}
	}
	return coord
}
*/
