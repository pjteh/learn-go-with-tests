package iteration

const repeatCount = 5

func Repeat(char string) string {
	returnString := ""
	for i := 0; i < repeatCount; i++ {
		returnString += char
	}
	return returnString
}
