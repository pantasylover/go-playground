My Golang sandbox that consists of test programs and such.

* [Go & C](./go_and_c/) - Programs that interact with C.
    * [Basic01](./go_and_c/basic01/)

## Building
You may use `make` to build all Go programs and libraries into the [`build/`](./build) directory. You may also use `make clean` to cleanup all build files.

## Notes
### Library Path With Shared C Libraries
For Go programs that utilize shared C libraries, you may need to add [`build/libs`](./build/libs/) to your `LD_LIBRARY_PATH` environmental variable.

## Credits
* [Christian Deacon](https://github.com/gamemann)