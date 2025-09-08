package main
import (
	"fmt"
)
	type order struct{
		id int
		items []int
		total float64
		adress string
		isCompleted bool
	} 
func main(){
	var a order = order{1,[]int{1,2},123.1,"a",true}
	var orderMap = map[int]order{}
	fmt.Println(a)
	fmt.Println(orderMap)
	orderMap = addOrder(a,orderMap)
	fmt.Println(orderMap)
}
func addOrder(a order,b map[int]order)(map[int]order){
	b[a.id]= a
	return b
}