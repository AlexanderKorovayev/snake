/*
package game
модуль transport
модуль для реализации транспортных функций между клиентом и сервером
*/

package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func getSnakeCoord() []coordinates {
	message := new(TransportData) //{"test", map[string][]coordinates{"test": {{1, 1}}}, "test1"}
	message.MainObjectsCoord = map[string][]coordinates{}
	//сообщим серверу имя клиента
	message.Info = getOutboundIP()
	// так как передаём инфу в виде набора байт
	bytesMessageRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(string(bytesMessageRepresentation))

	r := bytes.NewReader(bytesMessageRepresentation)
	resp, err := http.Post("http://localhost:8080/create", "application/json", r)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	// надо всегда закрывать боди иначе соединение не закроется
	defer resp.Body.Close()

	// приведём результат к заданной структуре
	var res TransportData
	err = json.Unmarshal(body, &res)

	if err != nil {
		//добавить обработку ошибок
	}

	//fmt.Println(res)
	// вытащим координаты из результата
	return parseSnakeCoord(&res)
}

// получить ip клиента, что бы сервер мог
// однозначно идентифицировать каждого клиента
func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func parseSnakeCoord(data *TransportData) []coordinates {
	// получаем координаты
	return data.MainObjectsCoord[getOutboundIP()]
	// обрабатываем сообщение
	//data.Info необходимо организовать обработку сообщений
}
