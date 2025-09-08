package main
import "fmt"
import "strings"
import "errors"
func main(){
	var name string = "Максим"
	var age int = 19
	var email  string = "maxim@mail"
	fmt.Println(validateUser(name,age,email))
}
func validateUser(name string, age int, email string)(error){
	if name=="" {
		return errors.New("Имя не должно быть пустым")
	}
	if len(name)>=50{
		return errors.New("Длина имени не должна превышать 50 символов")
	}
	if age<18||age>120{
		return errors.New("Диапозон возраст должен быть от 18 до 120")
	}
	if strings.Index(email,"@")==-1{
		return errors.New("Неправильный email")
	}
	return nil
}
