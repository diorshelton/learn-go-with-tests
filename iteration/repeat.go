package iteration

func Repeat(character string) string {
	var repeated = ""
	for i:=0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated   
}
