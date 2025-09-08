package main

import "fmt"

type InventoryItem struct {
	Name     string
	Weight   float64
	IsQuestItem bool
}
func main() {
	inventory := []InventoryItem{
		{"Экскалибур!!!", 3.5, false},
		{"Зелье лечения", 0.5, true},
		{"Щит", 5.0, false},
		{"Посох", 1.2, true},
		{"Лук", 2.0, false},
	}
	fmt.Println("Инвентарь:", inventory)
	totalWeight := calcWeight(inventory)
	fmt.Println("Общий вес инвентаря:", totalWeight)
}
func calcWeight(items []InventoryItem) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Weight
	}
	return total
}

