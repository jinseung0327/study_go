package pt


func Pot() int {
	total := 0
	v1 := 2
	v2 := &v1
	*v2 = 16
	total = v1
	return total
}