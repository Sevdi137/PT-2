package main
import "fmt"
type Employee struct{
	id int
	name string
	position string
	salary float64
}
func main(){
	emp := []Employee{{1,"Максим","Работяга", 2000}, {2,"Сергей","Тоже работяга", 1000}}
	fmt.Println(salaryAvg(emp))
}
func salaryAvg(a []Employee)(float64,float64){
	var sum float64 = 0
	for _,sal:= range a{
		sum += sal.salary
	}
	return sum,sum/float64(len(a))
}