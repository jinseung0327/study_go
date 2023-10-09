package main

import (
	"Gopractice/accounts"
	"Gopractice/mydict"
	"fmt"
)

// func lenAndUpper(name string) (lenght int, uppercase string) {
// 	defer fmt.Println("END")
// 	lenght = len(name);
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// type person struct {
// 	name string
// 	age int
// 	favFood []string
// }

func main() {
	account := accounts.NewAccount("jinseung")
	account.Deposit(10)
	fmt.Println(account)

	dictionary := mydict.Dictionary{}
	//baseWord := "hello"
	// definition := "Greeting"
	// err := dictionary.AddbaseWord, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.SearchbaseWord)
	// fmt.Println("found",baseWord, "definition:", hello)
	// err2 := dictionary.AddbaseWord, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)


	// dictionary := mydict.Dictionary{"first": "FirstbaseWord"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }


	// err := account.Withdraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// favFood := []string{"kimchi", "ramen"}
	// nico := person{name: "nico", age: 18, favFood: favFood}
	// fmt.Println(nico)


	// value := loof.SuperAdd(1, 2, 3, 4, 5)
	// value2 := loof.SuperAdd2(1, 2, 3, 4, 5)
	// value3 := pt.Pot()
	// fmt.Println(value)
	// fmt.Println(value2)
	// fmt.Println(tf.CanIDrink(15))
	// fmt.Println(tf.CanIDrink2(50))
	// fmt.Println(value3)
}