package mock

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, 3)
}
