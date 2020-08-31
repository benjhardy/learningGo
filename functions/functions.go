package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beSocial(name string) {
	sayHello(name)
	fmt.Println("Hows the weather?")
	sayGoodbye(name)
}

func addOne(x int) int {
	return x + 1
}

func sayHelloABunch() {
	fmt.Println("Hello")
	sayHelloABunch()
}
func main() {

	beSocial("Ben")
	beSocial("Bailey")

	x := 5
	x = addOne(x)
	fmt.Println(x)

	x = addOne(addOne(x))

	fmt.Println(x)

	sayHelloABunch()

}
