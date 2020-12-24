/*
package core
модуль area
отвечает за отричовку игрового поля.
небольшое описание для termloop:
	есть canvas(полотно) это двумерный массив, который содержит Cell(ячейки),
	они могут содержать разные символы и цвета, таким образом рисуется канвас.
	Что бы объект имел возможность попасть в канвас он должен отнаследоваться от
	энтити.
	основные модули это:
		entity.go
		input.go
		level.go
		screen.go
		termloop.go
*/

package game

import (
	"encoding/json"
	"time"

	"github.com/JoelOtter/termloop"
)

// флаг, сигнализирующий, нужно ли ещё ожидать игроков и начала основной игры
var initGameFlag bool = true

//CreateArea создать арену, по которой будет перемещаться змейка
func CreateArea() *area {
	area := new(area)
	area.Entity = termloop.NewEntity(1, 1, 1, 1)
	border := make(map[Coordinates]int)
	fillAreaBorder(width, high, &border)
	area.areaBorder = border
	return area
}

//fillAreaBorder заполнить игровую область информацией о её границах
func fillAreaBorder(imax, jmax int, border *map[Coordinates]int) {
	starti := 0
	startj := 0
	for i := starti; i < imax; i++ {
		for j := startj; j < jmax; j++ {
			coord := Coordinates{i, j}
			if i == starti || i == imax-1 {
				(*border)[coord] = 1
			} else if j == startj || j == jmax-1 {
				(*border)[coord] = 1
			} else {
				(*border)[coord] = 0
			}
		}
	}
}

//Draw отвечает за отрисовку змеи на дисплее
func (area *area) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	for k, v := range area.areaBorder {
		if v == 1 {
			screen.RenderCell(k.X, k.Y, &termloop.Cell{Fg: termloop.ColorWhite,
				Bg: termloop.ColorWhite})
		}
	}
}

//Tick отслеживаем события
func (area *area) Tick(event termloop.Event) {
	// если ожидание игроков ещё требуется, то опрашиваем сервер
	if initGameFlag {
		// далее опрашиваем сервер до тех пор пока не будут готовы все игроки
		// отправляем серваку свою готовность играть.
		// внутри себя сервак запускает обратный отсчёт на добавление
		// остальных игроков и будет ждать только это время.
		// по истечению этого времени он отошлёт сообщение о готовности играть
		// а также координаты для всех объектов.
		// клиент в бесконечном цикле опрашивает сервер и если в ответе число
		// то отрисовываем его, если в ответе ready, то рисуем все объекты и
		// дальше по тику делаем запросы на перерисовку
		// всех объектов и получаем координаты.

		// опрашиваем сервер
		logToFIle("start loop")
		info := getServerInfo()
		// распарсим info в json
		infoJSON := new(TransportData)
		infoJSON.MainObjectsCoord = map[string][]Coordinates{}
		err := json.Unmarshal(info, infoJSON)
		if err != nil {
			//добавить обработку ошибок
		}
		logToFIle(infoJSON)
		// теперь надо добавить проверку infoJSON на то что внутри.
		estimate := parseServerInfo(infoJSON)
		logToFIle(estimate)
		// отрисуем обратный отсчёт
		// мы уже создали глобальный GameScreen в startBaseSnakeLevel, поэтому тут
		// надо просто его дополнить
		GameScreen.TimeToReady = CreateTimeObj(estimate)
		GameScreen.AddEntity(GameScreen.TimeToReady)
		// добавим остальные объекты на уже созданный уровень
		//startMainSnakeLevel()
		time.Sleep(time.Second * 3)
		logToFIle("finish loop")
	}
}

//Collision произошло ли косание с змейкой
func (area *area) collision(c *Coordinates) bool {
	return area.areaBorder[*c] == 1
}
