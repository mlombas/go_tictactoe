package io

import "fmt"

func AskUser(what string) {
   fmt.Printf("Please enter %s: ", what)
}

func ReportUserMistake(message string) {
   fmt.Printf("Oops! it seems you made a mistake: %s\n", message)
   fmt.Printf("Please try again: ")
}
