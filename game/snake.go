package game

import (
	tl "github.com/JoelOtter/termloop"
)

// Player сущность человечка
// должна наследовать энтити что бы быть в игре
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

//Draw позволяет двигать камеру за персонажем
func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

// переменная, которая хранит направление змейки
var snakerout string = "Up"

// функция для определения движения змейки в тех случаях,
// когда никаких событий не происходит
func defaulttick(player *Player, event tl.Event) {
	player.prevX, player.prevY = player.Position()
	if snakerout == "Right" {
		player.SetPosition(player.prevX+1, player.prevY)
	}
	if snakerout == "Left" {
		player.SetPosition(player.prevX-1, player.prevY)
	}
	if snakerout == "Up" {
		player.SetPosition(player.prevX, player.prevY-1)
	}
	if snakerout == "Down" {
		player.SetPosition(player.prevX, player.prevY+1)
	}
}

// Tick позволяет сущности человека двигаться
// каждый момент времени мы проверяем состояние, что бы отрисовать объекты
func (player *Player) Tick(event tl.Event) {
	if event.Type != tl.EventKey {
		// движение по умолчанию
		defaulttick(player, event)
	} else {
		player.prevX, player.prevY = player.Position()
		if event.Key == tl.KeyArrowRight {
			player.SetPosition(player.prevX+1, player.prevY)
			snakerout = "Right"
		}
		if event.Key == tl.KeyArrowLeft {
			player.SetPosition(player.prevX-1, player.prevY)
			snakerout = "Left"
		}
		if event.Key == tl.KeyArrowUp {
			player.SetPosition(player.prevX, player.prevY-1)
			snakerout = "Up"
		}
		if event.Key == tl.KeyArrowDown {
			player.SetPosition(player.prevX, player.prevY+1)
			snakerout = "Down"
		}
	}
}

// Collide позволяет сущности человека взаимодействовать с другими объектами
func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}
