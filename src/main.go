package main

import (
   "./application"
   "./application/board"
   "./application/player"
   "./io"
)

func main() {
   user := player.NewUserPlayer()
   ai := player.NewAIPlayer(
      "HAHA! I got you!",
      "Mmmm, this is too easy",
      "There is absolutely no way you can win this",
      "Woah, I was not expecting that!",
      "So much options... So many alternatives...",
      "Aw come on, I wanted that tile",
      "You never give up, do you?",
   )

   io.Announce("Ancient master", "The fate of the world is in your hands now.")
   io.Announce("Ancient master", "You MUST defeat this foe.")
   io.Announce("Ancient master", "Otherwise humanity is doomed.")
   io.Announce("Ancient master", "Remember what I taught you.")
   io.Announce("Ancient master", "You are crosses! Now go and figth.")

   b := board.NewBoard()

   game := application.NewGame(user, ai, b)
   result := game.PlayUntilEnd()

   io.PrintStringable(&result) 
   winner := result.WhoWon()
   if winner == board.NOTHING {
      io.PrintDraw()
   } else {
      io.PrintVictory(winner.String() + " player")
   }
}
