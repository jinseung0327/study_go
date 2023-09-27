package loof

func SuperAdd(numbers ...int) int {
    sum := 0 

    for i := 0; i < len(numbers); i++ {
        sum += numbers[i] 
    }

    return sum 
}

func SuperAdd2(numbers2 ...int) int {
	total := 0
	for _, number := range numbers2 {
		total	+= number
	}
	return total
}
