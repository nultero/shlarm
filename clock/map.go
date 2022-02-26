package clock

var NumsMap = map[int][]string{
	0: zero,
	1: one,
	2: two,
	3: three,
	4: four,
	5: five,
	6: six,
	7: seven,
	8: eight,
	9: nine,
}

func LeftBuf() []string {
	b := []string{}
	for i := 0; i < Height; i++ {
		b = append(b, Space)
	}
	return b
}
