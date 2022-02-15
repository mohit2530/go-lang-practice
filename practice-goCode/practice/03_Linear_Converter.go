package practice

import (
	"fmt"
	"strings"
)

// lengthConverterInput - returns the user input
func lengthConverterInput() (float64, string) {

	fmt.Println("Please enter the length: ")
	var lengthToConvert float64
	_, err := fmt.Scanln(&lengthToConvert)

	if err != nil {
		return 0, "error: cannot convert the length"
	}

	fmt.Println("Is the measurement in (m)eters or (f)eet? ")
	var lengthCategory string
	_, err = fmt.Scanln(&lengthCategory)

	if err != nil {
		return 0, "error: cannot convert the length"
	}
	return lengthToConvert, lengthCategory

}

// ConvertLength - returns the converted length
func ConvertLength() string {

	lengthToConvert, lengthCategory := lengthConverterInput()

	if strings.EqualFold(lengthCategory, "f") {
		convertedLength := fmt.Sprintf("%.2f", lengthToConvert*0.3048)
		fmt.Printf("%ff is %sm", lengthToConvert, convertedLength)
		return convertedLength
	}
	convertedLength := fmt.Sprintf("%.2f", lengthToConvert*3.2808399)
	fmt.Printf("%fm is %sf", lengthToConvert, convertedLength)
	return convertedLength

}
