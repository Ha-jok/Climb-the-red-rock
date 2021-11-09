package main
import(
	"fmt"
	"time"
)
var (
	x,y  int
	c string
)
func main ()  {
	fmt.Println("intput:")
	fmt.Scan(&x,&c,&y)
	switch c {
	case "+" :
		fmt.Println(x,c,y,"=",x+y)
	case "-":
		fmt.Println(x,c,y,"=",x-y)
	case "*":
		fmt.Println(x,c,y,"=",x*y)
	case "/":
		fmt.Println(x,c,y,"=",float64(x)/float64(y))
	default:
		fmt.Println("Error")
	}
	time.Sleep(5*time.Second)

}