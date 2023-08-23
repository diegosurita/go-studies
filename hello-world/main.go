// Every Go program must start with a package declaration.
// The main package is a special package. It tells the Go compiler that the package should compile as an executable program instead of a shared library.
package main

// fmt is a standard library package.
// Standard library packages: https://pkg.go.dev/std
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
