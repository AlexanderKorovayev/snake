package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AlexanderKorovaev/snake/server/core"
)

//Coordinates координаты
type Coordinates struct {
	X int
	Y int
}

type TransportData struct {
	Action string
	Coord  map[string][]Coordinates
	Info   string
}

func main() {
	// надо передавать служебную инфу в json
	// надо понять, как правильно читать данные из боди

	/*
		message := map[string]interface{}{
			"hello": "world",
			"life":  42,
			"test": map[string]string{
				"yes": "of course!",
			},
		}
	*/
	message := core.TransportData{"test", map[string][]core.Coordinates{"test": {{1, 1}}}, "test1"}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytesRepresentation))

	r := bytes.NewReader(bytesRepresentation)
	resp, err := http.Post("http://localhost:8080/hello", "application/json", r)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	// надо всегда закрывать боди иначе соединение не закроется
	defer resp.Body.Close()

	fmt.Println(string(body))
}
