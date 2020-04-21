package application

import (
   "strings"
)

type Tile int
const (
   CROSS = -1
   NOTHING = 0
   NOTCH = 1
)

func isValidTile(t Tile) bool {
   return t == CROSS || t == NOTHING || t == NOTCH
}

func (t Tile) String() string {
   switch t {
   case CROSS: return "X"
   case NOTCH: return "O"
   default: return " "
   }
}

type Board struct {
   tiles [9]Tile
   turn Tile
}

func NewBoard() Board {
   return Board {
      turn: CROSS,
   }
}

func (b *Board) Turn(x, y int) {
   b.Set(x, y, b.turn)
   b.switchTurn()
}

func (b *Board) switchTurn() {
   if b.turn == CROSS {
      b.turn = NOTCH
   } else {
      b.turn = CROSS
   }
}

func (b *Board) set(x, y int, t Tile) {
   if isValidTile(t) {
      b.tiles[y * 3 + x] = t
   }
}

func (b *Board) get(x, y int) Tile {
   return b.tiles[y * 3 + x]
}

func (b *Board) IsWon() Tile {
   var res Tile
   for x := 0; x < 3; x++ {
      res = b.checkCol(x)
      if res != NOTHING {
         return res
      }
   }
   for y := 0; y < 3; y++ {
      res = b.checkRow(y)
      if res != NOTHING {
         return res
      }
   }

   if res = b.checkMainDiagonal(); res != NOTHING {
      return res
   }
   if res = b.checkSecondDiagonal(); res != NOTHING {
      return res
   }

   return NOTHING
}

func (b *Board) checkCol(x int) Tile {
   s := make([]Tile, 3)
   for y := 0; y < 3; y++ {
      s[y] = b.get(x, y)
   }

   return checkSlice(s)
}

func (b *Board) checkRow(y int) Tile {
   s := make([]Tile, 3)
   for x := 0; x < 3; x++ {
      s[x] = b.get(x, y)
   }

   return checkSlice(s)
}

func (b *Board) checkMainDiagonal() Tile {
   s := make([]Tile, 3)
   for i := 0; i < 3; i++ {
      s[i] = b.get(i, i)
   }

   return checkSlice(s)
}

func (b *Board) checkSecondDiagonal() Tile {
   s := make([]Tile, 3)
   for i := 0; i < 3; i++ {
      s[i] = b.get(i, 3 - 1 - i)
   }

   return checkSlice(s)
}

func checkSlice(s []Tile) Tile {
   count := 0
   for _, v := range s {
      count += int(v)
   }

   switch count {
   case 3*NOTCH: return NOTCH
   case 3*CROSS: return CROSS
   default: return NOTHING
   }
}

func (b *Board) String() string {
   lines := make([]string, 3)
   for y := 0; y < 3; y++ {
      line := make([]string, 3)
      for x := 0; x < 3; x++ {
         line[x] = b.get(x, y).String()
      }
      lines[y] = strings.Join(line, "|")
   }

   interline := "\n-----\n"
   return strings.Join(lines, interline)
}
