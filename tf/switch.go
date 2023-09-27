package tf

func CanIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge{
	case 18:
		return false
	case 17:
		return true
	case 50:
		return false
	}
	return false
}