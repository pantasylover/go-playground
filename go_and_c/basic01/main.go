package main

// #cgo LDFLAGS: -lbasic01
// #include "basic01.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	field1 := 10
	field2 := 45
	field3 := "This is a test string!"

	// We need to create a string using C.CString() and then free it later.
	cField3 := C.CString(field3)

	// Create structure.
	test := C.Test{
		field1: C.int(field1),
		field2: C.int(field2),
		field3: cField3,
	}

	fmt.Printf("C Test Stucture: field1 => %d, field2 => %d, field3 => '%s'.\n", test.field1, test.field2, C.GoString(test.field3))

	// Free field3 string.
	C.free(unsafe.Pointer(cField3))
}
