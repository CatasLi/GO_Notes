package main 
import "fmt"
import "time"
func main() {
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	switch{       //just like if/else
	case time.Now().Hour() < 12 :
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good afternoon!")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("I am a bool.")
		case int:
			fmt.Println("I am a integer.")
		default:
			fmt.Println("Unknown type:", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}