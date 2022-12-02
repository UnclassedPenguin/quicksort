package main

import (
  "fmt"
  "github.com/gdamore/tcell/v2"
)

func mainLoop(s tcell.Screen, style tcell.Style) {
  x, y := s.Size()

  writeToScreen(s, style, 1, 1, "Press 1 to sort")
}

// This is used just to write strings to the screen. Used in the "menu".
func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), []rune{}, style)
  }
}

func main() {
  s, err := tcell.NewScreen()
  if err != nil {
    fmt.Println("Error in tcell.NewScreen:", err)
  }

  if err := s.Init(); err != nil {
    fmt.Println("Error initializing screen:", err)
    os.Exit(1)
  }

  s.Clear()

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

  mainLoop(s, style)
}
