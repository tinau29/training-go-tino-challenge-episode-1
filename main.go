package main 

import (
   "log"
   "episode-satu/app"
)

func main() {
   // Task 1
   log.Println(app.Multiply(7.8, 2))

   // Task 2
   input := "1"
   if isValid, err := app.IsInteger(input); nil != err {
      log.Println(isValid, err.Error())
   }

   // Task 3
   inputDays := 3
   log.Println(app.IntToDayName(inputDays))

   // Task 4
   daysName := "jumat" 
   log.Println(app.DayNameToInt(daysName))

   // Task 5 
   var kosong string
   app.IsiHello(&kosong)
   log.Println(kosong)
}

