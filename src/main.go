package main

import (
   "./io"
   "fmt"
)

func main() {
   v := io.AskIntBounded("enter your age", 0, 100)
   fmt.Println(v)
}
