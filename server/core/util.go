/*
package core
модуль util
содержит вспомогательные функции
*/

package core

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"gopkg.in/ini.v1"
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

// Countdown функция для обратного отсчёта перед началом игры
func Countdown() {
	for TimeCount > 0 {
		//fmt.Println(TimeCount)
		TimeCount--
		time.Sleep(time.Second)
	}
}

//GetCoordinates получение рандомных координат
func GetCoordinates() (int, int) {
	// инициализируем рандомизатор для оси X
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// инициализируем рандомизатор для оси Y
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	x := r1.Intn(Width - 1)
	if x == 0 {
		x++
	}
	y := r2.Intn(High - 1)
	if y == 0 {
		y++
	}
	return x, y
}

// Remove удаление из среза
func Remove(s []string, el string) []string {
	for i, name := range s {
		if el == name {
			s[i] = s[len(s)-1]
		}
	}
	return s[:len(s)-1]
}

// ParseBody приведение данных от клиента к нужному виду
func ParseBody(r *http.Request) *TransportData {
	// если все условия соблюдены, то начинаем читать данные
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//добавить обработку ошибок
	}

	// приводим данные к нужном формату
	data := CreateTransportData()
	err = json.Unmarshal(body, &data)

	if err != nil {
		//добавить обработку ошибок
	}
	return data
}

func readConfig(path string) *ini.File {
	cfg, err := ini.Load(path)
	if err != nil {
		// сделать обработчик ошибок
	}
	return cfg
}

// InitializationGlobals инициализация глобальных переменных
func InitializationGlobals(path string) {
	conf := readConfig(path)
	Width, _ = conf.Section("").Key("width").Int()
	High, _ = conf.Section("").Key("high").Int()
	MaxObjectsCount, _ = conf.Section("").Key("MaxObjectsCount").Int()
	TimeCount, _ = conf.Section("").Key("TimeCount").Int()
	Colors = conf.Section("").Key("Colors").Strings(",")
}
