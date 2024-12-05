package main

// #cgo pkg-config: python-3.10-embed
// #include <Python.h>
// int CPyList_Check(PyObject* o) {
//     return PyList_Check(o);
// }
//
// int CPyFloat_Check(PyObject* o) {
//     return PyFloat_Check(o);
// }
import "C"
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)
 
func main() {

    fmt.Printf("Hello from go\n")
    pycodeGo := `
import sys
for path in sys.path:
    print(path)
` 
    // pycodeGo := `
// print("Hello from python")
    // `
   
    defer C.Py_Finalize()
    C.Py_Initialize()
    if C.Py_IsInitialized() == 0 {
        fmt.Println("Error initializing the python interpreter")
        os.Exit(1)
    }

    dir, err := filepath.Abs("./")
    if err != nil {
        log.Fatal(err)
    }

    importString := C.CString("import sys\nsys.path.append(\"" + dir + "\")")
    defer C.free(unsafe.Pointer(importString))
    ret := C.PyRun_SimpleString(importString)

    if ret != 0 {
        log.Fatalf("error appending '%s' to python sys.path", dir)
    }

    pycodeC := C.CString(pycodeGo)
    defer C.free(unsafe.Pointer(pycodeC))
    C.PyRun_SimpleString(pycodeC)

    import_lib_str := C.CString("kuma")
    defer C.free(unsafe.Pointer(import_lib_str))
    oImport := C.PyImport_ImportModule(import_lib_str)
    if !(oImport != nil && C.PyErr_Occurred() == nil) {
        fmt.Printf("hello\n")
        C.PyErr_Print()
        log.Fatal("failed to import kuma")
    }
    defer C.Py_DecRef(oImport)

    str := C.CString("kuma")
    defer C.free(unsafe.Pointer(str))
    oModule := C.PyImport_AddModule(str)
    if !(oModule != nil && C.PyErr_Occurred() == nil) {
        C.PyErr_Print()
        log.Fatal("failed to add module kuma")
    }
    // defer oModule.DecRef()
    

    demo(oModule)

 
}

func demo(module *C.PyObject) {

    testdata, err := getTestData(module)
    if err != nil {
        log.Fatalf("Error gettign testdata: %s", err)
    }
    fmt.Printf("test data: %#v\n", testdata)
}

func getTestData(module *C.PyObject) ([]float64, error) {
    str := C.CString("gen_test_data") 
    defer C.free(unsafe.Pointer(str))
    pFunc := C.PyObject_GetAttrString(module, str)
    if pFunc == nil || C.PyCallable_Check(pFunc) == 0 {
        C.PyErr_Print()
        return nil, fmt.Errorf("Could not get gen_test_data")
    }
    defer C.Py_DecRef(pFunc)

    pArgs := C.PyTuple_New(0)
    pValue := C.PyObject_CallObject(pFunc, pArgs)
    C.Py_DecRef(pArgs)
    if pValue == nil {
        C.PyErr_Print()
        return nil, fmt.Errorf("Could not get return value for testdata")
    }
    defer C.Py_DecRef(pValue)

    goArr, err := goSliceFromPyList(pValue)
    if err != nil {
        return nil, err
    }

    return goArr, nil
}

func python_object_get_type(obj *C.PyObject) string {
    repr := C.PyObject_Repr(C.PyObject_Type(obj))

    // Convert the PyObject to a C string
    cStr := C.PyUnicode_AsUTF8(repr)

    // Convert the C string to a Go string
    goStr := C.GoString(cStr)
    return goStr

}

func goSliceFromPyList(pyList *C.PyObject) ([]float64, error) {
    if C.CPyList_Check(pyList) == 0 {
        goStr := python_object_get_type(pyList)
        return nil, fmt.Errorf("Object given is not of type list, type: %v", goStr)
    }
    slice := make([]float64, 0)
    size := int(C.PyList_Size(pyList))
    for i := range size {
        item := C.PyList_GetItem(pyList, C.Py_ssize_t(i))
        if C.CPyFloat_Check(item) == 0 {
            goStr := python_object_get_type(item)
            return nil, fmt.Errorf("Object is not of type float, type: %v", goStr)
        }

        value := float64(C.PyFloat_AsDouble(item))
        if C.PyErr_Occurred() != nil {
            C.PyErr_Print()
            return nil, fmt.Errorf("Could not get float64 from python float")
        }
        slice = append(slice, value)

    } 
    return slice, nil
}
