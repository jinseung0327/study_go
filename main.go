package main

import (
	"Gopractice/loof"
	"Gopractice/pt"
	"Gopractice/tf"
	"fmt"
)

// func lenAndUpper(name string) (lenght int, uppercase string) {
// 	defer fmt.Println("END")
// 	lenght = len(name);
// 	uppercase = strings.ToUpper(name)
// 	return
// }


func main() {
	value := loof.SuperAdd(1, 2, 3, 4, 5)
	value2 := loof.SuperAdd2(1, 2, 3, 4, 5)
	value3 := pt.Pot()
	fmt.Println(value)
	fmt.Println(value2)
	fmt.Println(tf.CanIDrink(15))
	fmt.Println(tf.CanIDrink2(50))
	fmt.Println(value3)
}