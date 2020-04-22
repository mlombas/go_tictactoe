package main

import (
   "./application"
   "./application/player"
   "fmt"
)

func main() {
   u := player.NewUserPlayer()
   b := application.NewBoard()
   for !b.Ended() {
      fmt.Println(&b)
      announceTurn(b.GetTurn())
      
      x, y := u.GetMove(b)
      b.PlayTurn(x, y)
   }
   fmt.Println(&b)
   showWinner(b.WhoWon())
}

func showWinner(t application.Tile) {
   if t == application.CROSS {
      fmt.Println("Cross won!")
   } else {
      fmt.Println("Notch won!")
   }
}

func announceTurn(t application.Tile) {
   if t == application.CROSS {
      fmt.Println("Its cross player turn!")
   } else {
      fmt.Println("Its notch player turn!")
   }
}
