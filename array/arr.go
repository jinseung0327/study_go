package array

import "fmt"

func Arr() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}