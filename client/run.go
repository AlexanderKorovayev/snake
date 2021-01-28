//package входная точка для запуска игры
package main

import (
	"github.com/AlexanderKorovaev/snake/client/game"
)

func main() {
	// сохраним необходимые настройки из конфига
	game.InitializationGlobals("config.ini")
	game.StartGame()
}
