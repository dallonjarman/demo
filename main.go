package main
// CGO CGLAGS: -fplugin=./attack.so
// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
// 	return f();
// }
//
// int fourtyTwo()
// {
// 	return 42;
// }
import "C"
import "fmt"

func main() {
	f := C.intFunc(C.fourtytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 42
}
