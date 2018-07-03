package main

import "C"

//export sum
func sum(a, b C.int) C.int {
   return a + b
}

func main() {
// This is necessary for the compiler.
// You can add something that will be executed when engaging your library to the interpreter.
}
