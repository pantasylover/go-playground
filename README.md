My Golang playground that consists of test programs and such.

* [Channels](./chans/) - Go programs that interact with channels.
    * [Chan01](./chans/chan01/) - A Go program that reads and writes to a string channel in separate Go routines.
* [Go & C](./go_and_c/) - Go Programs that interface with C via the [Cgo](https://pkg.go.dev/cmd/cgo) package using structures, functions, and more!
    * [Basic01](./go_and_c/basic01/) - A Go program that creates and prints a basic C structure (integer and string fields).
    * [Basic02](./go_and_c/basic02/) - A Go program that creates a *person* at the age of 26 and then increases their age to *27* 10 seconds later using C structures and functions.
    * [Pointer01](./go_and_c/pointer01) - A Go program that creates and retrieves values from a C structure with different types of integer fields/sizes and a string. Also uses functions to retrieve/set fields and allocates the C structure with `malloc()` so you need to free it inside of the main Go program.
    * [Array01](./go_and_c/array01/) - A Go program that passes *people* created from a config file to a C function that dynamically allocates memory for each person and returns a separate array of the people passed to the function.
    * [Array02](./go_and_c/array02/) - A Go program that does the exact same thing from Array01. However, in this version, the way dynamic entries are allocated is different and allows you to specify the max entries when retrieving from the Go program via `GetPeople()`.
    * [Dlsym01](./go_and_c/dlsym01/) - Loads C functions from a shared object at runtime and executes the functions.
* [Structures](./structs/) - Go programs that interact with structures.
    * [InterfaceToStruct01](./structs/interface_to_slice01/) - A Go program that converts/casts a field that has the `interface{}` type to a custom structure.

## Building
You may use `make` to build all Go programs and libraries into the [`build/`](./build) directory. You may also use `make clean` to cleanup all build files.

## Notes
### Library Path With Shared C Libraries
For Go programs that utilize shared C libraries, you may need to add [`build/libs`](./build/libs/) to your `LD_LIBRARY_PATH` environmental variable.

## Credits
* [Christian Deacon](https://github.com/gamemann)