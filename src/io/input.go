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
         ReportUserMistake("input was not an int")
      } else {
         return int(result)
      }
   }
}

func GetIntBounded(min, max int) value int {
   for !util.IsInBound(value, min, max) {
      ReportUserMistake("value out of bounds")
      value = GetInt()
   }

   return
}

func AskInt(message string) int {
   AskUser(message)
   return GetInt()
}

func AskIntBounded(message string, min, max int) int {
   AskUser(message)
   return GetIntBounded(min, max)
}

func AskString(message string) string {
   AskUser(message)
   return GetString()
}
