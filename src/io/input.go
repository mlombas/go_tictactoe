package io

import (
   "../util"
   "bufio"
   "errors"
   "os"
   "strconv"
   "strings"
)

func GetString() string {
   stdin := bufio.NewReader(os.Stdin)
   s, _ := stdin.ReadString('\n')
   s = strings.Trim(s, " \r\n")
   return s
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

func GetIntBounded(min, max int) (value int) {
   value = GetInt()
   for !util.IsInBound(value, min, max) {
      ReportUserMistake("value out of bounds")
      value = GetInt()
   }

   return
}

func GetIntPair() (int, int) {
   for {
      s := GetString()
      a, b, err := extractPair(s)
      if err != nil {
         ReportUserMistake("entered bad pair")
      } else {
         return a, b
      }
   }
}

func extractPair(s string) (int, int, error) {
   parts := strings.Split(s, ",")
   if len(parts) < 2 { return 0, 0, errors.New("Tuple too short") }
   aStr := strings.Trim(parts[0], " ")
   bStr := strings.Trim(parts[1], " ")

   a, err := strconv.ParseInt(aStr, 0, 64)
   if err != nil { return 0, 0, err }
   b, err := strconv.ParseInt(bStr, 0, 64)
   if err != nil { return 0, 0, err }

   return int(a), int(b), nil
}

func GetIntPairBounded(aMin, aMax, bMin, bMax int) (int, int) {
   a, b := GetIntPair()
   for !(util.IsInBound(a, aMin, aMax) && util.IsInBound(b, bMin, bMax)) {
      ReportUserMistake("values entered out of bounds")
      a, b = GetIntPair()
   }

   return a, b
}

func GetOfOptions(opts ...byte) byte {
   v := GetString()[0]
   for !util.ContainsByte(opts, v) {
      ReportUserMistake("value not in options")
      v = GetString()[0]
   }

   return v
}

func AskInt(message string) int {
   AskUser(message)
   return GetInt()
}

func AskIntBounded(message string, min, max int) int {
   AskUser(message)
   return GetIntBounded(min, max)
}

func AskIntPair(message string) (int, int) {
   AskUser(message)
   return GetIntPair()
}

func AskIntPairBounded(message string, aMin, aMax, bMin, bMax int) (int, int) {
   AskUser(message)
   return GetIntPairBounded(aMin, aMax, bMin, bMax)
}

func AskString(message string) string {
   AskUser(message)
   return GetString()
}

func AskOfOptions(message string, opts ...byte) byte {
   AskUser(message)
   return GetOfOptions(opts...)
}
