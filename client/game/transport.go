/*
package game
модуль transport
модуль для реализации транспортных функций между клиентом и сервером
*/

package game

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func getServerInfo() []byte {
	message := new(TransportData)
	message.MainObjectsCoord = map[string][]Coordinates{}
	// сообщим серверу имя клиента
	message.Info = getOutboundIP()
	// передаём инфу в виде набора байт
	bytesMessageRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewReader(bytesMessageRepresentation)
	resp, err := http.Post("http://localhost:8080/create", "application/json", r)
	if err != nil {
		//добавить обработку ошибки
	}

	body, _ := ioutil.ReadAll(resp.Body)

	// надо всегда закрывать боди иначе соединение не закроется
	defer resp.Body.Close()
	return body
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

func parseSnakeCoord(data *TransportData) []Coordinates {
	// получаем координаты
	return data.MainObjectsCoord[getOutboundIP()]
	// обрабатываем сообщение
	//data.Info необходимо организовать обработку сообщений
}
