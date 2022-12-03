package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)
func menu(s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  strings := []string{ "Unclassed Penguin Quick Sort",
                       "Press 1 to start from random seed",
                       "(You can also press 1 at any time",
                       "while it is running to restart",
                       "with a new seed.)",
                       "Esc, Ctrl-C, or q to quit",
                     }

  // Write strings to screen.
  for i, str := range strings {
    writeToScreen(s,style,((x/2)-(len(str)/2)),y/3+(i*2),str)
  }

  // Keyboard handling. Keys to quit (Esc, Ctrl-c, q)
  // and the key to start the game (1)
  for {
    switch ev := s.PollEvent().(type) {
    case *tcell.EventResize:
      s.Sync()
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC, tcell.KeyEscape:
        s.Fini()
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case 'q', 'Q':
          s.Fini()
          os.Exit(0)
        case '1':
          slice := createRandomSlice(s)
          mainLoop(slice, s, style)
        }
      }
    }
  }
}


func mainLoop(slice [][]int, s tcell.Screen, style tcell.Style) {
  x, y := s.Size()

  go func() {
    for {
      switch ev := s.PollEvent().(type) {
      case *tcell.EventResize:
        s.Sync()
      case *tcell.EventKey:
        switch ev.Key() {
        case tcell.KeyCtrlC, tcell.KeyEscape:
          s.Fini()
          os.Exit(0)
        case tcell.KeyRune:
          switch ev.Rune() {
          case 'q', 'Q':
            s.Fini()
            os.Exit(0)
          case '1':
            s.Clear()
            draw(slice, s, style)
            writeToScreen(s, style, 1, 1, "Press 1 For new array")
            writeToScreen(s, style, 1, 2, "Press 2 to sort")
            writeToScreen(s, style, 1, 3, fmt.Sprintf("X: %v", x))
            writeToScreen(s, style, 1, 4, fmt.Sprintf("Y: %v", y))
            s.Sync()
          case '2':
            s.Clear()
            writeToScreen(s, style, 1, x/2, "SORTING...")
            s.Sync()
          }
        }
      }
    }
  }()
}

func draw(slice [][]int, s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  s.Clear()
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      if slice[i][j] == 1 {
        s.SetContent(i, j, tcell.RuneBlock, []rune{}, style)
      }
    }
  }
  s.Sync()

}

func createRandomSlice(s tcell.Screen) [][]int {
  x, y := s.Size()

  var slice [][]int

  rand.Seed(time.Now().UnixNano())

  for i := 0; i < x; i++ {
    var newSlice []int
    slice = append(slice, newSlice)
    rand := rand.Intn(y)
    for j := 0; j < y; j++ {
      var newInt int
      if j < rand {
        newInt = 0
      } else {
        newInt = 1
      }

      slice[i] = append(slice[i], newInt)
    }
  }
  return slice
}

// This is used just to write strings to the screen. Used in the "menu".
func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), []rune{}, style)
  }
}

func flipCoin(total int, limit int) bool{
  x := rand.Intn(total)
  if x <= limit {
    return true
  } else {
    return false
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

  menu(s, style)
}
