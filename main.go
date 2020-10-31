package main

//#cgo CFLAGS: -x objective-c
//#cgo LDFLAGS: -framework Foundation
import "C"
import (
	"fmt"

	"github.com/diegomagdaleno/tagger/converters"
)

func main() {
	thing := converters.GetFinalArrayOfFiles("/Users/me/Documents/Screenshots")
	fmt.Println("No idea ")
	fmt.Println(thing)
}
