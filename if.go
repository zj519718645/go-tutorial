package main 
import "fmt"
import "math"

func sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprint(math.Sqrt(-x))
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2),sqrt(-4))
}