package main 
import "fmt"
import "time"
func main() {
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:      //multi-conditions separated by comma
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	switch{                              //just like if/else
	case time.Now().Hour() < 12 :
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good afternoon!")
	}

	whatAmI := func(i interface{}){     //type switch compares types instead of values
		switch t := i.(type){
		case bool:
			fmt.Println("I am a bool.")
		case int:
			fmt.Println("I am a integer.")
		default:
			fmt.Println(t)
			fmt.Printf("Unknown type: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}