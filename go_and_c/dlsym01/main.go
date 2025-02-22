package main

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>

int callFunc1(void *f) {
    int (*fn)() = (int (*)())f;
    return fn();
}

void callFunc2(void *f, int x, int y) {
    void (*fn)(int, int) = (void (*)(int, int))f;
    fn(x, y);
}

int callFunc3(void *f, char *x, int y, int z) {
    int (*fn)(const char*, int, int) = (int (*)(const char*, int, int))f;

	return fn(x, y, z);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	libPath := C.CString("./build/libs/libdlsym01.so")
	defer C.free(unsafe.Pointer(libPath))

	hdl := C.dlopen(libPath, C.RTLD_LAZY)

	if hdl == nil {
		fmt.Println("Failed to open libdlsym01.so...")

		return
	}

	defer C.dlclose(hdl)

	func1SymName := C.CString("Func1")
	defer C.free(unsafe.Pointer(func1SymName))

	func1Sym := C.dlsym(hdl, func1SymName)

	if func1Sym == nil {
		fmt.Println("Failed to load 'Func1' symbol...")

		return
	}

	func2SymName := C.CString("Func2")
	defer C.free(unsafe.Pointer(func2SymName))

	func2Sym := C.dlsym(hdl, func2SymName)

	if func2Sym == nil {
		fmt.Println("Failed to load 'Func2' symbol...")

		return
	}

	func3SymName := C.CString("Func3")
	defer C.free(unsafe.Pointer(func3SymName))

	func3Sym := C.dlsym(hdl, func3SymName)

	if func3Sym == nil {
		fmt.Println("Failed to load 'Func3' symbol...")

		return
	}

	fmt.Println("Running 'Func1'...")

	ret := C.callFunc1(func1Sym)

	fmt.Printf("Return code => %d\n", ret)

	fmt.Println("Running 'Func2'...")

	C.callFunc2(func2Sym, C.int(5), C.int(30))

	fmt.Println("Running 'Func3'...")

	xVal := C.CString("hello world")
	defer C.free(unsafe.Pointer(xVal))

	ret = C.callFunc3(func3Sym, xVal, C.int(13), C.int(50))

	fmt.Printf("Return code => %d\n", ret)
}
