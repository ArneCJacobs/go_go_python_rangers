package main

// #cgo pkg-config: python-3.10-embed
// #include <Python.h>
import "C"
import (
	"fmt"
	"unsafe"
)
 
func main() {

    fmt.Printf("Hello from go\n")
    pycodeGo := `
import sys
print("Hello from python")
for path in sys.path:
    print(path)
` 
   
    defer C.Py_Finalize()
    C.Py_Initialize()
    pycodeC := C.CString(pycodeGo)
    defer C.free(unsafe.Pointer(pycodeC))
    C.PyRun_SimpleString(pycodeC)
 
}
