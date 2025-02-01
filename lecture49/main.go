package main

import (
	"fmt"

	"github.com/KevMonstah/Puppy"
)

func main() {
	s1 := Puppy.Bark()
	//s2 := Puppy.Barks()
	fmt.Println(s1)
	//fmt.Println(s2)
	fmt.Println(Puppy.Barks())

	fmt.Println(Puppy.Sleeps("Happy"))
	fmt.Println(Puppy.Runs("Gilmore"))
	fmt.Println(Puppy.Chills("Boo-yah"))

}
