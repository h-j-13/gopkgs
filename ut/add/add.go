package add

func Add(a, b interface{}) (c interface{}) {
	va, ok := a.(int)
	if !ok {
		return 0
	}
	vb, ok := b.(int)
	if !ok {
		return 0
	}

	return va + vb
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
