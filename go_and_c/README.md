Go Programs that interface with C via the [Cgo](https://pkg.go.dev/cmd/cgo) package using structures, functions, and more!

* [Basic01](./basic01/) - A Go program that creates and prints a basic C structure (integer and string fields).
* [Basic02](./basic02/) - A Go program that creates a *person* at the age of 26 and then increases their age to *27* 10 seconds later using C structures and functions.
* [Pointer01](./pointer01) - A Go program that creates and retrieves values from a C structure with different types of integer fields/sizes and a string. Also uses functions to retrieve/set fields and allocates the C structure with `malloc()` so you need to free it inside of the main Go program.
* [Array01](./array01/) - A Go program that passes *people* created from a config file to a C function that dynamically allocates memory for each person and returns a separate array of the people passed to the function.
* [Array02](./array02/) - A Go program that does the exact same thing from Array01. However, in this version, the way dynamic entries are allocated is different and allows you to specify the max entries when retrieving from the Go program via `GetPeople()`.