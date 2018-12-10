package main
import "fmt"
import "math"

const s string = "constant"
func main() {
	fmt.Println(s)
	const n = 500000
	const d = 3e20 / n
	fmt.Println(d)             //numeric constants have no type
	fmt.Println(int64(d))      //they will have a type when given(eg. explicit cast)
	fmt.Println(math.Sin(n))   //numeric constants get type from context(eg. Sin(float64))
}