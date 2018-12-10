package main 
import "fmt"
func main() {
	s := make([]string, 3) //use make to create a slice, no [len]
	fmt.Println("s:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s:", s)
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "f")
	fmt.Println("apd s:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	l := s[2:4]    //slice operator !![low,high)!!   b[:] = b ; x := b[:] x is a slice from array b 
	fmt.Println("s[2:4]:", l)

	l = s[:5]
	fmt.Println("s[:5]:", l)

	l = s[2:]
	fmt.Println("s[2:]:", l) 

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i + 1) //lengthes vary
		for j := 0; j < i + 1; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}