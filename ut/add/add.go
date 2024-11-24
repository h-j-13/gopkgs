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
