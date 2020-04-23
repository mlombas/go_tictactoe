package board

import (
   "strings"
)

type Position struct {
   x int
   y int
}
func NewPosition(x, y int) Position {
   return Position { x, y }
}
func (pos Position) Extract() (int, int) {
   return pos.x, pos.y
}

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

func (b *Board) Copy() Board {
   var newTiles [9]Tile
   for i, v := range b.tiles {
      newTiles[i] = v
   }

   return Board {
      tiles: newTiles,
      turn: b.turn,
   }
}

func (b *Board) GetTurn() Tile {
   return b.turn
}

//This is here for future extensibility
func (b *Board) GetDimension() int {
   return 3
}

func (b *Board) PlayTurn(pos Position) bool {
   if b.get(pos) == NOTHING {
      b.set(pos, b.turn)
      b.switchTurn()

      return true
   } else {
      return false
   }
}

func (b *Board) switchTurn() {
   if b.turn == CROSS {
      b.turn = NOTCH
   } else {
      b.turn = CROSS
   }
}

func (b *Board) set(pos Position, t Tile) {
   if isValidTile(t) {
      x, y := pos.Extract()
      b.tiles[y * 3 + x] = t
   }
}

func (b *Board) get(pos Position) Tile {
   x, y := pos.Extract()
   return b.tiles[y * 3 + x]
}

func (b *Board) IsOccupied(pos Position) bool {
   return b.get(pos) != NOTHING
}

func (b *Board) Ended() bool {
   return b.GetNFreeTiles() == 0 || b.WhoWon() != NOTHING
}

func (b *Board) GetNFreeTiles() (counter int) {
   for _, v := range b.tiles {
      if v == NOTHING {
         counter++
      }
   }

   return
}

func (b *Board) WhoWon() (res Tile) {
   for x := 0; x < 3; x++ {
      res = b.checkCol(x)
      if res != NOTHING {
         return
      }
   }
   for y := 0; y < 3; y++ {
      res = b.checkRow(y)
      if res != NOTHING {
         return
      }
   }

   if res = b.checkMainDiagonal(); res != NOTHING {
      return
   }
   if res = b.checkSecondDiagonal(); res != NOTHING {
      return
   }

   return
}

func (b *Board) checkCol(x int) Tile {
   s := make([]Tile, 3)
   for y := 0; y < 3; y++ {
      s[y] = b.get(NewPosition(x, y))
   }

   return checkSlice(s)
}

func (b *Board) checkRow(y int) Tile {
   s := make([]Tile, 3)
   for x := 0; x < 3; x++ {
      s[x] = b.get(NewPosition(x, y))
   }

   return checkSlice(s)
}

func (b *Board) checkMainDiagonal() Tile {
   s := make([]Tile, 3)
   for i := 0; i < 3; i++ {
      s[i] = b.get(NewPosition(i, i))
   }

   return checkSlice(s)
}

func (b *Board) checkSecondDiagonal() Tile {
   s := make([]Tile, 3)
   for i := 0; i < 3; i++ {
      s[i] = b.get(NewPosition(i, 3 - 1 - i))
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
         line[x] = b.get(NewPosition(x, y)).String()
      }
      lines[y] = strings.Join(line, "|")
   }

   interline := "\n-----\n"
   return strings.Join(lines, interline)
}
