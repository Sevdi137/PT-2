package main

import (
	"fmt"
)

func main() {
	tracks := map[string]float64{
		"Еда":        15000,
		"Транспорт":  5000,
		"Развлечение": 3000,
	}
	tracks["Еда"] += 2000
	total := 0.0
	for _, value := range tracks {
		total += value
	}
	fmt.Println("Итоговая сумма:", total)
}
