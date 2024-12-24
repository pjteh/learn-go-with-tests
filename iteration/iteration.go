package iteration

var repeatCount = 0

func Repeat(char string, times int) string {
	repeatCount = times
	returnString := ""
	for i := 0; i < repeatCount; i++ {
		returnString += char
	}
	return returnString
}
