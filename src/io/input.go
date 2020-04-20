package io

import (
   "../util"
   "fmt"
   "strconv"
)

func GetString() string {
   var text string
   fmt.Scan(&text)
   return text
}

func GetInt() int {
   for {
      input := GetString()
      result, err := strconv.ParseInt(input, 0, 64)

      if err != nil {
         UserMistake("input was not an int")
      } else {
         return int(result)
      }
   }
}

func AskInt(message string) int {
   AskUser(message)
   return GetInt()
}

func AskIntBounded(message string, min, max int) int {
   AskUser(message)

   value := GetInt()
   for !util.IsInBound(value, min, max) {
      UserMistake("value out of bounds")
      value = GetInt()
   }

   return value
}

func AskString(message string) string {
   AskUser(message)
   return GetString()
}
