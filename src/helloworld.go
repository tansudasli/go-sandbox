// executable (main) vs reusable
package main

import "fmt"

// go build -o out src/helloworld.go
// to run
//  ./out/helloworld           or
//  go run src/helloworld.go
func main() {
	fmt.Println("Hi, there")
}
