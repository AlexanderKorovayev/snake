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
	"fmt"

	"github.com/JoelOtter/termloop"
)

// флаг, сигнализирующий, нужно ли ещё ожидать игроков и начала основной игры
var initGameFlag bool = true

//createArea создать арену, по которой будет перемещаться змейка
func createArea() *area {
	area := new(area)
	area.Entity = termloop.NewEntity(1, 1, 1, 1)
	border := make(map[coordinates]int)
	fillAreaBorder(width, high, &border)
	area.areaBorder = border
	return area
}

//fillAreaBorder заполнить игровую область информацией о её границах
func fillAreaBorder(imax, jmax int, border *map[coordinates]int) {
	starti := 0
	startj := 0
	for i := starti; i < imax; i++ {
		for j := startj; j < jmax; j++ {
			coord := coordinates{i, j}
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

		// создадим сообщение, которое необходимо передать серверу
		message := createTransportData()
		// опрашиваем сервер
		info := getServerInfo("initiate", message)
		// распарсим info в json
		err := json.Unmarshal(info, message)
		if err != nil {
			//добавить обработку ошибок
		}
		// получим статус
		status := message.Info
		// по статусу определяем сценарий действий, определяемый в свиче
		switch status {
		case "added", "already added":
			// значит нам пришёл обратный отсчёт
			estimate := message.Estimate
			// отрисуем обратный отсчёт
			// мы уже создали глобальный GameScreen в startBaseSnakeLevel, поэтому тут
			// надо просто обновлять в нём обратный отсчёт
			gameScreen.timeToReady = createTimeObj(fmt.Sprintf("your color is %v. Start in %v", message.Color[clientID], estimate))
			gameScreen.AddEntity(gameScreen.timeToReady)
		case "busy":
			// реализовать обработку
		case "finished":
			// реализовать обработку
		case "ready":
			// отключаем отрисовку чисел
			// хз но RemoveEntity не работает,
			// при этом если добавлять пустой объект без пробела,
			// то он не обновляется и остаётся последнее
			// добавленное число, поэтому оставил пробел.Но нужно заменять все объекты
			gameScreen.timeToReady = createTimeObj("                                  ")
			gameScreen.AddEntity(gameScreen.timeToReady)
			// отключим инициализирующий игру цикл
			// теперь они будут исходить от тика змейки у каждого игрока
			initGameFlag = false
			// получим направления для змееки
			// infoJSON.Info имеет тип интерфейс, но если в нём хранится сложный
			// объект, то его поля тоже имеют тип интерфейс, поэтому будет
			// преобразовывать его поэтапно
			directionMap := message.DirectionMap
			// добавим остальные объекты на уже созданный уровень
			startMainSnakeLevel(message.MainObjectsCoord, directionMap, message.Color)
		}
	}
}

//Collision произошло ли косание с змейкой
func (area *area) collision(c *coordinates) bool {
	return area.areaBorder[*c] == 1
}
