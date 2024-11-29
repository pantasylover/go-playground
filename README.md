My Golang playground that consists of test programs and such.

* [Go & C](./go_and_c/) - Go Programs that interface with C via the [Cgo](https://pkg.go.dev/cmd/cgo) package using structures, functions, and more!
    * [Basic01](./go_and_c/basic01/) - A Go program that creates and prints a basic C structure (integer and string fields).
    * [Basic02](./go_and_c/basic02/) - A Go program that creates a *person* at the age of 26 and then increases their age to *27* 10 seconds later using C structures and functions.

## Building
You may use `make` to build all Go programs and libraries into the [`build/`](./build) directory. You may also use `make clean` to cleanup all build files.

## Notes
### Library Path With Shared C Libraries
For Go programs that utilize shared C libraries, you may need to add [`build/libs`](./build/libs/) to your `LD_LIBRARY_PATH` environmental variable.

## Credits
* [Christian Deacon](https://github.com/gamemann)