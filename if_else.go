package main

import "fmt"

func main() {
	if 7 % 2 == 0 {
		fmt.Println("7 is even.")
	}else{
		fmt.Println("7 is odd.")
	}
	if num := 9; num < 0 {                  //声明可以先与条件，且在其他分支上也可使用
		fmt.Println("num is negative.")
	}else if num < 10 {
		fmt.Println("num has one digit.")
	}else{
		fmt.Println("num has many digits.")
	}
}
