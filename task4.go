package main

import (
	"fmt"
)
func main(){
	var names []string = []string{"Анна","Борис","Виктор"}
	var votes map[string]string = map[string]string{"Матвей":string(names[0]),"Николай":(names[1]),"Андрей":(names[1]), "Алейксей": (names[1]),"Никита":(names[2]),"Семён":(names[2]),}
voteCounting(votes)
}
func voteCounting(votes map[string]string){
	sum := float64(len(votes))
	results := map[string]int{}
	for _, vote := range votes{
		results[vote]++
	}
	for name, count := range results{
		fmt.Println(name,count, float64(count)/float64(sum)*100,"%")
	}
}