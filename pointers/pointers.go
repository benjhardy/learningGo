package main

import "fmt"

func addOne(num *int) {
	*num = *num + 1
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) {
	// when using pointers to structs, Go doesnt require
	// * notation for each struct element
	x := b.pos.x
	y := b.pos.y
	// guy ->pos.x would be used if guy wasnt a struct
	fmt.Println("(", x, ",", y, ")")
}

func main() {

	x := 5
	fmt.Println(x)

	var xPtr *int = &x

	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

	p := position{4, 2}
	badguy := badGuy{"Jabba", 100, p}

	whereIsBadGuy(&badguy)

}
