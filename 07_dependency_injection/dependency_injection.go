package dependencyinjection

import (
	"fmt"
	"io"
)

// with DI we specify a Writer interface so our function handles only data creation,
// but doesn't handle where the data is written to (separation of concerns)
// this way we can write data to different outputs (stdout, bytes.Buffer, http.ResponseWriter)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
