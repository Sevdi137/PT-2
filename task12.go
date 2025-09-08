package main

import (
	"fmt"
	"strconv"
)

const (
	bin = 2
	dec = 10
	hex = 16
)
func main() {
	var number int64
	fmt.Print("Введите десятичное число: ")
	fmt.Scanln(&number)
	convert(number)
}
func convert(num int64) {
	fmt.Println("Десятичное:", strconv.FormatInt(num, dec))
	fmt.Println("Двоичное:", strconv.FormatInt(num, bin))
	fmt.Println("Шестнадцатеричное:", strconv.FormatInt(num, hex))
}
