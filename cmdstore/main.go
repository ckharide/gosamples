package main
import (
	"fmt"

	"github.com/ckharide/cmdstore/mathpkg"
	"github.com/ckharide/cmdstore/stringpkg"
)

func main() {
	fmt.Println(mathpkg.ComputeAdd(10, 20))
	mathpkg.GetValues()
	a := "Hello"
	b := "world"
	a, b = stringpkg.Swap("Hello", "World")
	fmt.Println(a)
	fmt.Printf(b)
}
