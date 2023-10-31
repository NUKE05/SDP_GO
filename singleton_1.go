package main

import (
	"fmt"
	"sync"
)

type GameSettings struct {
	difficulty   string
	soundEnabled bool
}

var instance *GameSettings
var once sync.Once

func GetGameSettings() *GameSettings {
	once.Do(func() {
		instance = &GameSettings{
			difficulty:   "Легко",
			soundEnabled: true,
		}
	})
	return instance
}

func (gs *GameSettings) SetDifficulty(difficulty string) {
	gs.difficulty = difficulty
}

func (gs *GameSettings) EnableSound() {
	gs.soundEnabled = true
}

func (gs *GameSettings) DisableSound() {
	gs.soundEnabled = false
}

func (gs *GameSettings) GetSettings() string {
	return fmt.Sprintf("Сложность: %s, Звук отключен: %v", gs.difficulty, gs.soundEnabled)
}

func main() {
	// Отображение изначальных настроек
	settings := GetGameSettings()
	fmt.Println(settings.GetSettings())
	// Процесс изменения настроек и дальнейшее отображение результата
	settings.SetDifficulty("Сложно")
	settings.DisableSound()
	fmt.Println(settings.GetSettings())

	otherSettings := GetGameSettings()
	fmt.Println(otherSettings.GetSettings())

	otherSettings.SetDifficulty("Средний")
	settings.EnableSound()
	fmt.Println(settings.GetSettings())
}
