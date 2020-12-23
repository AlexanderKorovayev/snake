/*
package core
модуль util
содержит вспомогательные функции
*/

package game

import (
	"log"
	"os"
)

//FindInSlice функция для поиска вхождений в срезе
func FindInSlice(sliceData *[]Coordinates, data *Coordinates) bool {
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
