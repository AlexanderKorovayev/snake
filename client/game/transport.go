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

func getServerInfo(postMethodName string, message *transportData) []byte {
	// сообщим серверу имя клиента
	message.ClientID = clientID
	// передаём инфу в виде набора байт
	bytesMessageRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewReader(bytesMessageRepresentation)
	resp, err := http.Post(fmt.Sprintf("http://localhost:2000/%v", postMethodName), "application/json", r)
	//resp, err := http.Post(fmt.Sprintf("https://wet-horse-80.loca.lt/%v", postMethodName), "application/json", r)
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

func parseSnakeCoord(data *transportData) []coordinates {
	// получаем координаты
	return data.MainObjectsCoord[clientID]
	// обрабатываем сообщение
	//data.Info необходимо организовать обработку сообщений
}
