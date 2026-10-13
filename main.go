# main
package main

import "fmt"

func HelloCodeCup(n int) string {
	return fmt.Sprintf("Hello CodeCup %d", n)
}

func main() {
	fmt.Println(HelloCodeCup(7))
}
