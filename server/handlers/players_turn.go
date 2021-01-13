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
	"math/rand"
	"net/http"
	"time"

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
	clName := (data.Info).(string)
	core.MainObjects[clName] = data.MainObjectsCoord[clName]
	// просчитаем для данного клиента новые координаты
	// сначала приводим интерфейс к типу флоат а потом его уже к типу дирекшен
	drctn := core.Direction((data.Action).(float64))
	core.MainObjects[clName] = updateSnakeCoordinates(core.MainObjects[clName], drctn)
	// запишем клиенту все координаты
	data.MainObjectsCoord = core.MainObjects
	// определим, было ли столкновение с едой
	if foodCollision(core.MainObjects[clName]) {
		// увеличиваем змейку
		core.MainObjects[clName] = increaseSnake(core.MainObjects[clName], drctn)
		// перемещаем еду на новое место
		x, y := getCoordinates()
		core.MainObjects["food"] = []core.Coordinates{{X: x, Y: y}}
	}
	// обратно посылаем статусы, поэтому клиент будет ожидать строку
	data.Action = ""
	data.Info = ""
	// определим, было ли столкновение змейки с самой собой
	if snakeSelfCollision(core.MainObjects[clName]) {
		// отправим статус о коллисии
		data.Action = "snakeSelfCollision"
		// !!!!!так же необходимо удалить унфу об игроке!!!!!
	}
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

// получение головы змейки
func getHead(body []core.Coordinates) core.Coordinates {
	return body[len(body)-1]
}

// определение коллизии с едой
func foodCollision(body []core.Coordinates) bool {
	snakeHead := getHead(body)
	return core.MainObjects["food"][0].X == snakeHead.X && core.MainObjects["food"][0].Y == snakeHead.Y
}

// Где считать врезание с окружением?
//areaCollision определение коллизии с окружением
/*
func (snake *snake) areaCollision() bool {
	return GameScreen.GameArea.collision(snake.GetHead())
}
*/

// определение столкновений змейки с самой собой
func snakeSelfCollision(body []core.Coordinates) bool {
	bodyWithoutHead := body[:len(body)-1]
	snakeHead := getHead(body)
	return core.FindInSlice(&bodyWithoutHead, &snakeHead)
}

// увеличить змейку
func increaseSnake(body []core.Coordinates, drctn core.Direction) []core.Coordinates {
	if drctn == core.Right {
		head := body[len(body)-1]
		head.X++
		body = append(body, head)
	}
	if drctn == core.Left {
		head := body[len(body)-1]
		head.X--
		body = append(body, head)
	}
	if drctn == core.Up {
		head := body[len(body)-1]
		head.Y--
		body = append(body, head)
	}
	if drctn == core.Down {
		head := body[len(body)-1]
		head.Y++
		body = append(body, head)
	}
	return body
}

// получение рандомных координат для пищи
func getCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси Y
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	x := r1.Intn(core.Width - 1)
	if x == 0 {
		x++
	}
	y := r2.Intn(core.High - 1)
	if y == 0 {
		y++
	}
	return x, y
}
