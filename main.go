package main

import (
	"fmt"
)

func main() {
	Drivers := map[string]string{
		"Charles LeClerc":    "Ferarri",
		"Carlos Sainz":       "Ferarri",
		"Yuki Tsunoda":       "Alpha Tauri",
		"Pierre Gasly":       "Alpha Tauri",
		"Mick Schumacher":    "Hass",
		"Nikita Mazepin":     "Hass",
		"George Russell":     "Williams",
		"Nicholas Latifi":    "Williams",
		"Antonio Giovinazzi": "Alfa Romeo Racing",
		"Kimi Raikkonen":     "Alfa Romeo Racing",
		"Lance Stroll":       "Aston Martin",
		"Sebastian Vettal":   "Aston Matin",
		"Esteban Ocon":       "Alpine",
		"Fernando Alonso":    "Alpine",
		"Lando Norris":       "McLaren",
		"Daniel Ricciardo":   "McLaren",
		"Lewis Hamilton":     "Mercedes",
		"Valtteri Bottas":    "Mercedes",
		"Sergio Perez":       "Red Bull Racing",
		"Max Verstappen":     "Red Bull Racing",
	}

	printMap(Drivers)
}

func printMap(c map[string]string) {
	for Drivers, team := range c {
		fmt.Println(Drivers, "Drives for", team)
	}
}
