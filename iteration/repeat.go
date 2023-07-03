package iteration
import "fmt"

func Repeat(character string, repeatCount int) string {
	var repeat string
	for i := 0; i < repeatCount; i++ {
			repeat = repeat + character
	}
  return repeat
}

func ExampleRepeat() {
	repeat := Repeat("a", 9)
	fmt.Printf(repeat)
	// Output: "aaaaaaaaa"
}