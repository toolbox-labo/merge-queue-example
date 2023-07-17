package main

import "fmt"

func main() {
	fmt.Printf("%s", greeting("John"))
}

func getMSG() string {
	return "Hi"
}

func greeting(name string) string {
	return fmt.Sprintf("%s, %s", getMSG(), name)
}
