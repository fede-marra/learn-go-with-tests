package iteration

func Repeat(character string, cant int) string {
	var repeated string
	for i := 0; i<cant; i++ {
		repeated += character
	}
	return repeated
}