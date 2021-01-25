/*
package server
когда нажали ентер клиент посылает сигнал серверу о том, что один игрок подключился.
Тем временем на клиенте создаётся поле, но змеек ещё нет. Ждём ответ от сервера.
Сервер принимает подключение, добавляет его в свой пул подключений и начинает ждать 30сек что бы подключился хотя бы
ещё один чел.
Тем временем сервер считает координаты размещения первого игрока и посылает их, так же посылает число обратного отсчёта.
После того как подключились два чела и прошло 30 сек сервер рассылает координаты всех объектов каждому клиенту.
Каждый тик мы отправляем направление змейки на сервер, он их отрабатывает.
Каждый тик мы принимаем координаты от сервера и отрисовываем их.
*/

package main

import (
	"net/http"

	"github.com/AlexanderKorovaev/snake/server/core"
	"github.com/AlexanderKorovaev/snake/server/handlers"
)

// подумать
// жлбавить конфиг файлы

func main() {
	// запустим обратный отсчёт
	// в целом тут просто надо, что бы посчитался отсчёт и
	// сообщил клиенту, что больше подключаться нельзя.
	// т.е. при следующих подключениях к /initiate он будет сообщать, что время вышло.
	// таким образом сервер будет предохраняться от новых подключений во время игры
	go core.Countdown()
	http.HandleFunc("/initiate", handlers.InitiateGame)
	http.HandleFunc("/playersTurn", handlers.PlayersTurn)
	http.ListenAndServe(":2000", nil)
}
