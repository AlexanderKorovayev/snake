/*
package game
модуль start_level
отвечает за отрисовку стартового меню.
*/

package game

import (
	"encoding/json"
	"time"

	"github.com/JoelOtter/termloop"
)

//startLevel стартовый уровень игры
type startLevel struct {
	termloop.Level
	startMenu *startMenu
}

//startMenu объект стартовое окно
type startMenu struct {
	*termloop.Text
}

//createStartMenu стартовое состояние игры
func createStartMenu() *startMenu {
	startObj := new(startMenu)
	startObj.Text = termloop.NewText(0, 0, "press Enter to srart game",
		termloop.ColorWhite,
		termloop.ColorDefault)
	return startObj
}

//Tick отслеживаем события
func (s *startMenu) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		// если нажали enter, то начинаем этап подготовки всех играков к игре
		if event.Key == termloop.KeyEnter {
			// сначала установим базовую часть уровня
			level := startBaseSnakeLevel()
			TermloopGame.Screen().SetLevel(level)
			time.Sleep(time.Second)
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

			// цикл опроса готовности сервера к игре
			// тут запускается бесконечный цикл, поэтому тик не заканчивается
			// и отрисовки не происходит
			for {
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
				time.Sleep(time.Second * 3)
				logToFIle("finish loop")
				/*
					// теперь надо добавить проверку infoJSON на то что внутри.
					estimate := parseServerInfo(infoJSON)
					// отрисуем обратный отсчёт
					level.TimeToReady = CreateTimeObj(estimate)
					level.AddEntity(GameScreen.TimeToReady)
					// добавим остальные объекты на уровень
					//startMainSnakeLevel()
					time.Sleep(time.Second * 3)
				*/
			}
		}
	}
}
