package main

import "fmt"

func main() {
	s := `package main

import "fmt"

func main() {
	s := 
	fmt.Print(s[:48] + "\u0060" + s + "\u0060" + s[48:])
}
`
	fmt.Print(s[:48] + "\u0060" + s + "\u0060" + s[48:])
}
