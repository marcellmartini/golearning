package stdout

import (
	"fmt"
	"io"
	"os"

	"basic/tools"
)

// Examples of stdout
func Examples() {
	tools.HeaderOutPut("pointerExamples")

	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground\n")
}
