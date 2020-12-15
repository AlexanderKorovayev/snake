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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AlexanderKorovaev/snake/server/core"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// надо научиться принимать и отправлять запросы и научиться их разбирать
	fmt.Printf("%v\n", r.Header)

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	//пробую создать необходимую
	//testData := core.TransportData{"test", map[string][]core.Coordinates{"test": {{1, 1}}}, "test1"}
	//fmt.Println(testData)
	var data core.TransportData
	// можно делать универсальные подходы при парсинге json
	// var testResp interface{}
	err := json.Unmarshal(body, &data)

	if err != nil {

	}

	// тип интерфейс, но к нему нельзя обращаться по индексу, поэтому необходимо
	// привести явно к типу
	// тоже самое придётся сделать и со значениями по ключу, потому что они
	// тоже типа интерфейс
	//testRespTyped := testResp.(map[string]interface{})
	//fmt.Println(testRespTyped)

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	/*
		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}
	*/

	data.Action = "alex555"
	fmt.Println(data)

	myJSON, _ := json.Marshal(data)
	fmt.Fprintf(w, string(myJSON))
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
