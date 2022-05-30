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
	emptyString := ""  //"input angka tidak sesuai "
	days := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	if number > 0  {
		numberDays := number - 1

		if numberDays >= 0 && numberDays < 7 {
			return days[numberDays]
		}

		return emptyString
	}
	
	return emptyString

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

	return days[daysNameLower]
	
}

func IsiHello(input *string) {
	*input = "hello world"
}


