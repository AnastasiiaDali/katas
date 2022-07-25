package align

func Align(str string, length int, alignment string) string {
	lenOfString := len(str)
	numberOfDots := length - lenOfString

	var finalStr string
	finalStr = str

	switch alignment {
	case "left":
		for i := 0; i < numberOfDots; i++ {
			finalStr = finalStr + "."
		}
	case "right":
		for i := 0; i < numberOfDots; i++ {
			finalStr = "." + finalStr
		}
	case "center":
		for i := 0; i < numberOfDots; i++ {
			if i == 0 || i%2 == 0 {
				finalStr = "." + finalStr
			} else {
				finalStr = finalStr + "."
			}
		}
	}
	return finalStr
}
