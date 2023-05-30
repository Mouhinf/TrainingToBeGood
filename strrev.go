package piscine

func StrRev(s string) string {
	var n int
	for index := range s {
		n = index
	}
	Word := []byte(s)
	b := 0
	for i := n; i >= 0; i-- {
		Word[i] = byte(s[b])
		b++
	}
	Final := string(Word)
	return Final
}
