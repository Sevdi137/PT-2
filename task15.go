package main

import (
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genres []string
}
func main() {
	movies := []Movie{
		{"Брат", 1997, 8.5, []string{"Драма", "Криминал"}},
		{"Левиафан", 2014, 7.8, []string{"Драма"}},
		{"Сталкер", 1979, 8.1, []string{"Фантастика", "Драма"}},
		{"Москва слезам не верит", 1980, 8.3, []string{"Комедия", "Драма"}},
		{"12", 2007, 8.0, []string{"Драма"}},		
	}
	fmt.Println("Фильмы:", movies)
	bestMovie := highestRated(movies)
	fmt.Println("Лучший фильм:",bestMovie.Title, bestMovie.Year,"- Рейтинг: ",bestMovie.Rating)}
func highestRated(movies []Movie) Movie {
	var best Movie
	for _, movie := range movies {
		if movie.Rating > best.Rating {
			best = movie
		}
	}
	return best
}
