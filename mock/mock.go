package mock

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	const finalWord = "Go!"
	const countdownStart = 3
	for i := countdownStart; i > 0; i-- {
		fmt.Fprint(out, i)
	}
	fmt.Fprint(out, finalWord)
}
