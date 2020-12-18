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

	"github.com/AlexanderKorovaev/snake/server/handlers"
)

func main() {
	// инициализируем глобальную переменну для
	// хранения колличества подключённых клиентов
	ClientsCount = make(map[string]string)

	http.HandleFunc("/initiate", handlers.GetSnakeCoordHandler)
	http.HandleFunc("/create", handlers.GetSnakeCoordHandler)
	http.HandleFunc("/getCoordinate", handlers.GetSnakeCoordHandler)
	http.ListenAndServe("localhost:8080", nil)
}
