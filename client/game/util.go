/*
package core
модуль util
содержит вспомогательные функции
*/

package game

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JoelOtter/termloop"
)

//FindInSlice функция для поиска вхождений в срезе
func FindInSlice(sliceData *[]coordinates, data *coordinates) bool {
	for _, el := range *sliceData {
		if el == *data {
			return true
		}
	}
	return false
}

func logToFIle(data interface{}) {
	// OpenFile принимает только один флаг.
	// перечисляя через OR мы можем обеспечить выполнение каждого флага.
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println(data)
	defer file.Close()
}

// удаление из среза
func remove(s []string, el string) []string {
	for i, name := range s {
		if el == name {
			s[i] = s[len(s)-1]
		}
	}
	return s[:len(s)-1]
}

// Соответствие цветов
func colorMap(color string) termloop.Attr {
	var colorTerm termloop.Attr
	if color == "Blue" {
		colorTerm = termloop.ColorBlue
	}
	if color == "Yellow" {
		colorTerm = termloop.ColorYellow
	}
	if color == "Green" {
		colorTerm = termloop.ColorGreen
	}
	if color == "Cyan" {
		colorTerm = termloop.ColorCyan
	}
	return colorTerm
}

// приведение данных от клиента к нужному виду
func parseBody(data_byte []byte) *transportData {
	// приводим данные к нужном формату
	data := createTransportData()
	err := json.Unmarshal(data_byte, &data)

	if err != nil {
		//добавить обработку ошибок
	}
	return data
}
