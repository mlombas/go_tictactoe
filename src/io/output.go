package io

import "fmt"

func Print(s string) {
   fmt.Println(s)
}

func AskUser(what string) {
   fmt.Printf("Please enter %s: ", what)
}

func UserMistake(message string) {
   fmt.Printf("Oops! it seems you made a mistake: %s\n", message)
   fmt.Printf("Please try again: ")
}
