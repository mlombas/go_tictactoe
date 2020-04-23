package io

import "fmt"

type Stringable interface {
   String() string
}

func AskUser(what string) {
   fmt.Printf("%s: ", what)
}

func ReportUserMistake(message string) {
   fmt.Printf("Oops! it seems you made a mistake: %s\n", message)
}

func AskUserToTryAgain(message string) {
   ReportUserMistake(message)
   fmt.Printf("Please try again: ")
}

func Announce(who string, message string) {
   fmt.Printf("%s: %s\n", who, message)
}

func PrintStringable(obj Stringable) {
   fmt.Println(obj.String())
}

func PrintVictory(winnerName string) {
   fmt.Printf("%s won! Congratulations!\n", winnerName)
}

func PrintDraw() {
   fmt.Println("Its a draw!")
}
