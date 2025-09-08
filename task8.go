package main
import "fmt"
type LogEntry struct{
	ip string
	httpId int
}
func main(){
	a := []LogEntry {{"192.168.1.1",404	},{"192.168.1.2",300},{"192.168.1.42",501}}
	fmt.Println(correctHttpId(a))
}
func correctHttpId(a []LogEntry)([]LogEntry){
	for i, log := range a{
		if log.httpId>600 || log.httpId<400{
			a = append(a[:i],a[i+1:]...)
		}
	}
	return a
}