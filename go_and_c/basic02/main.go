package main

// #cgo LDFLAGS: -lbasic02
// #include "basic02.h"
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	// Create a new person.
	firstName := "Christian"
	lastName := "Deacon"
	age := 26

	cFirstName := C.CString(firstName)
	cLastName := C.CString(lastName)
	cAge := C.int(age)

	person := C.CreatePerson(cFirstName, cLastName, cAge)

	// Print new structure returned.
	fmt.Printf("Created a person named '%s %s' who is %d years old!\n", C.GoString(person.firstName), C.GoString(person.lastName), person.age)

	// Wait 10 seconds.
	time.Sleep(10 * time.Second)

	// Change age to 27 using C function.
	C.ChangeAge((*C.Person)(&person), 27)

	fmt.Printf("Changed %s's age to %d!\n", C.GoString(person.firstName), person.age)

	// Free first and last name C strings.
	C.free(unsafe.Pointer(cFirstName))
	C.free(unsafe.Pointer(cLastName))
}
