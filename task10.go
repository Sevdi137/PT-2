package main
import ("fmt"; "strings")
type stats struct{
	symbols int
	words int
	sentences int
}
func main(){
	var text string = "Первоe предолжение. Второе предложение! Третье предложение?"
	stats := textStats(text)
	fmt.Println(text)
	fmt.Println(stats)
}
func textStats(text string)(stats){
	var stat stats
	stat.symbols=len(text)
	stat.words=len(strings.Fields(text))
	stat.sentences = 0
	for _,c := range text{
		if c == '.'|| c == '!' || c == '?'{
		stat.sentences++}
	}
	return stat
}

