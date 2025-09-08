package main

import "fmt"

type Product struct {
	Name     string
	Category string
	Price    float64
}

func main() {
	products := []Product{
		{"Морковь", "Овощи", 10.0},
		{"Картофель", "Овощи", 20.0},
		{"Майонез", "Соусы", 15.0},
	}
	fmt.Println("Изначальный список:",products)
	result := filterProducts(products, 15.0, "Овощи")
	fmt.Println("Отсортированный список:",result)
}
func filterProducts(products []Product, maxPrice float64, category string) []Product {
	var filtered []Product
	for _, p := range products {
		if p.Price < maxPrice && p.Category == category {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

