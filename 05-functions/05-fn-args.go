package main

import "fmt"

// decider that decides "which" function to execute
func main() {
	exec(f1) //=> execute f1
	exec(f2) //=> execute f2
	exec(func() {
		fmt.Println("anon fn invoked")
	})
}

// executor - executes the given function
func exec(fnRef func()) {
	fnRef()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
