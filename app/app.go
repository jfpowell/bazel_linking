package main

/*
#include "shared_library.h"
*/
import "C"

import "fmt"


func main() {
	fmt.Printf("main\n")
	fmt.Printf("C.function = %d\n", C.function())
}
