package main

import (
	"fmt"
)
func main(){
	var a,b,c float64
	fmt.Println("Введите вес основного багажа")
	fmt.Scanln(&a)
	fmt.Println("Введите вес ручной клади")	
	fmt.Scanln(&b)
	fmt.Println("Введите вес дополнительной ручной клади")	
	fmt.Scanln(&c)
	fmt.Println("Итоговый вес",a+b+c)	
}