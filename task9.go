package main

import "fmt"

const (
	single     = "single"
	double     = "double"
	suite      = "suite"
	free       = "free"
	booked     = "booked"
	maintenance = "maintenance"
)

type HotelRoom struct {
	Type   string
	Status string
	Price  float64
}

var rooms = make(map[string]HotelRoom)

func main() {
	rooms["101"] = HotelRoom{Type: single, Status: free, Price: 1500.0}
	fmt.Println(rooms["101"])
	bookRoom("101")
	fmt.Println(rooms["101"])
}
func bookRoom(roomNumber string) {
	if room, exists := rooms[roomNumber]; exists {
		room.Status = booked
		rooms[roomNumber] = room
	}
}
