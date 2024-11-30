package main

// #cgo LDFLAGS: -larray01
// #include "array01.h"
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/gamemann/go-playground/go_and_c/array01/config"
)

func GetPeople(people []C.Person, length int) []config.Person {
	if length < 1 {
		return nil
	}

	cPeople := &people[0]

	// Get People manually allocated memory.
	cPeopleRes := C.GetPeople(cPeople, C.int(length))

	if cPeopleRes == nil {
		return nil
	}

	defer C.FreePeople(cPeopleRes)

	var ret []config.Person

	for i := 0; i < int(cPeopleRes.length); i++ {
		ptr := cPeopleRes.entries[i]

		// Make sure we have a valid person pointer.
		if unsafe.Pointer(ptr) == unsafe.Pointer(nil) {
			fmt.Printf("[ERROR] GetPeople() :: Person at index %d is NULL.\n", i)

			continue
		}

		// Append to result.
		ret = append(ret, config.Person{
			FirstName: C.GoString(&ptr.firstName[0]),
			LastName:  C.GoString(&ptr.lastName[0]),
			Age:       int(ptr.age),
		})
	}

	return ret
}

func main() {
	// Create config and set defaults.
	cfg := config.Config{}
	cfg.SetDefaults()

	// Load from file system.
	err := cfg.Load("./conf.json")

	if err != nil {
		fmt.Printf("Error loading config from file system (./conf.json): %v\n", err)
	}

	// Create person array from C.
	cPeople := make([]C.Person, len(cfg.People))
	pplLen := 0

	// Loop through each person and create a separate person.
	for k, v := range cfg.People {
		cFirstName := C.CString(v.FirstName)
		defer C.free(unsafe.Pointer(cFirstName))

		cLastName := C.CString(v.LastName)
		defer C.free(unsafe.Pointer(cLastName))

		cAge := C.int(v.Age)

		cPeople[k] = C.CreatePerson(cFirstName, cLastName, cAge)

		p := &cPeople[k]

		fmt.Printf("Adding person: %s %s (%d years old)\n", C.GoString(&p.firstName[0]), C.GoString(&p.lastName[0]), p.age)

		pplLen++
	}

	if pplLen < 1 {
		fmt.Printf("Not enough people added from config.\n")

		os.Exit(1)
	}

	peopleFromPtr := GetPeople(cPeople, pplLen)

	fmt.Printf("Got %d people from GetPeople()\n", len(peopleFromPtr))

	if len(peopleFromPtr) > 0 {
		fmt.Println("People From GetPeople()")
		for _, v := range peopleFromPtr {
			fmt.Printf("\t- %s %s (Age %d)\n", v.FirstName, v.LastName, v.Age)
		}
	}
}
