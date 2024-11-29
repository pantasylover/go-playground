package main

// #cgo LDFLAGS: -lpointer01
// #include "pointer01.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Test fields.
	field1 := byte(155)
	field2 := uint16(52302)
	field3 := uint32(1231234)
	field4 := uint64(123512351232)
	field5 := "A test string"

	// Cast fields to C.
	cField1 := C.uchar(field1)
	cField2 := C.__u16(field2)
	cField3 := C.__u32(field3)
	cField4 := C.__u64(field4)
	cField5 := C.CString(field5)

	// Create test structure (returns a pointer).
	test := C.CreateTest(cField1, cField2, cField3, cField4, cField5)

	fmt.Printf("Created a test structure: field1 (u8) => %d, field2 (u16) => %d, field3 (u32) => %d, field4 (u64) => %d, field5 => %s\n", test.field1, test.field2, uint32(test.field3), uint64(test.field4), C.GoString(&test.field5[0]))

	// Free field5 string now.
	C.free(unsafe.Pointer(cField5))

	// Change field5 field using SetField5() function.
	field5 = "A new test string"
	cField5 = C.CString(field5)

	C.SetField5(test, cField5)

	// Free field5 string again.
	C.free(unsafe.Pointer(cField5))

	// Now retrieve field5 using GetField5() function and print it.
	field5Ret := C.GoString(C.GetField5(test))

	fmt.Printf("Changed test structure's field5 to: %s\n", field5Ret)

	// Now set field4 using SetField4().
	field4 = uint64(111122223334444)
	cField4 = C.__u64(field4)

	C.SetField4(test, cField4)

	// Now retrieve field4 using GetField4().
	field4Ret := uint64(C.GetField4(test))

	fmt.Printf("Changed test structure's field4 to: %d\n", uint64(field4Ret))

	// Free test structure since it was allocated via `malloc()`.
	C.free(unsafe.Pointer(test))
}
