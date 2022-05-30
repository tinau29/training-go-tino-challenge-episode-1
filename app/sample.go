package app 

import (
	"fmt"
	"errors"
	"reflect"
	"strings"
)

func Multiply(a float32 , b float32) (float32) {
	return a*b
}

func IsInteger(unknown interface{}) (bool, error) {
	var typeData interface{} = reflect.TypeOf(unknown)
	typeDataInput := fmt.Sprintf("%v", typeData)
	typeDataValid  := "int"
	if typeDataInput == typeDataValid {
		return true, nil 
		
	}

	return false, errors.New("Bukan Data Integer")
}

func IntToDayName(number int) (string) {
	days := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	if number > 0  {
		numberDays := number - 1

		if numberDays >= 0 && numberDays < 7 {
			return days[numberDays]
		}

		return "input angka tidak sesuai "
	}
	
	return "input angka tidak sesuai"

}

func DayNameToInt(daysName string ) (int) {
	daysNameLower := strings.ToLower(daysName)
	days := map[string]int {
		"senin": 1,
		"selasa": 2,
		"rabu": 3,
		"kamis": 4,
		"jumat": 5,
		"sabtu": 6,
		"minggu": 7,
	}

	switch(daysNameLower) {
		case "senin":
			return days[daysNameLower]
		case "selasa":
			return days[daysNameLower]
		case "rabu":
			return days[daysNameLower]
		case "kamis":
			return days[daysNameLower]
		case "jumat":
			return days[daysNameLower]
		case "sabtu":
			return days[daysNameLower]
		case "minggu":
			return days[daysNameLower]
		default:
			return 0
  }

	
}

// func IsiHello(input string) (string) {
// 	var typeData interface{} = reflect.TypeOf(input)
// 	typeDataInput := fmt.Sprintf("%v", typeData)
// 	return typeDataInput
// 	// check
// 	// if input  {
// 	// 	return "hello world 1"
// 	// }

// 	// return "hello world"

// }


