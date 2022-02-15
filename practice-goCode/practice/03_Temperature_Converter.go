package practice

import (
	"fmt"
	"strings"
)

/*
 	Temperature Converter

	Please enter the temperature : 58
	Is the temperature in (C)elsius to (F)ahrenheit : F
	Example : 58 F is 14 C
*/

//getUserInput - returns the user input
func getUserInput() (float64, string) {

	fmt.Println("Temperature to convert : ")
	var tempToConvert float64
	_, err := fmt.Scanln(&tempToConvert)
	if err != nil {
		return 0, "error : cannot convert the temperature"
	}

	fmt.Println("Is the temperature in (C)elsius to (F)ahrenheit :")
	var tempCategory string
	_, err = fmt.Scanln(&tempCategory)
	if err != nil {
		return 0, "error : cannot convert the temperature"
	}

	return tempToConvert, tempCategory
}

// ConvertTemperature - converts the temperature to the required items
func ConvertTemperature() string {

	tempToConvert, tempCategory := getUserInput()

	if strings.EqualFold(tempCategory, "F") {
		convertedTemp := fmt.Sprintf("%.2f", (tempToConvert-32)/1.8)
		fmt.Printf("%fF is %sC", tempToConvert, convertedTemp)
		return convertedTemp
	}
	convertedTemp := fmt.Sprintf("%.2f", (tempToConvert*1.8)+32)
	fmt.Printf("%fC is %sF", tempToConvert, convertedTemp)
	return convertedTemp

}
