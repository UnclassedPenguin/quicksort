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
  var slice [][]int

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
          s.Clear()
          slice = createRandomSlice(s)
          draw(slice, s, style)
          writeToScreen(s, style, 1, 1, "Press 1 For new random array")
          writeToScreen(s, style, 1, 2, "Press 2 For new \"ordered\" array")
          writeToScreen(s, style, 1, 3, "Press 3 to sort")
          writeToScreen(s, style, 1, 4, fmt.Sprintf("X: %v", x))
          writeToScreen(s, style, 1, 5, fmt.Sprintf("Y: %v", y))
          s.Sync()
        case '2':
          s.Clear()
          slice = createOrderedSlice(s)
          draw(slice, s, style)
          writeToScreen(s, style, 1, 1, "Press 1 For new random array")
          writeToScreen(s, style, 1, 2, "Press 2 For new \"ordered\" array")
          writeToScreen(s, style, 1, 3, "Press 3 to sort")
          writeToScreen(s, style, 1, 4, fmt.Sprintf("X: %v", x))
          writeToScreen(s, style, 1, 5, fmt.Sprintf("Y: %v", y))
          s.Sync()

        case '3':
          quickSort(s, style, slice, 0, len(slice)-1)
        }
      }
    }
  }
}

func countLength(slice []int, s tcell.Screen) int {
  _, y := s.Size()

  var count int

  for i := 0; i < y; i++ {
    if slice[i] == 1 {
      count++
    }
  }

  return count
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

func createOrderedSlice(s tcell.Screen, d int) [][]int {
  x, y := s.Size()

  var slice [][]int
  count := 1

  rand.Seed(time.Now().UnixNano())

  for i := 0; i < x; i++ {
    var newSlice []int
    slice = append(slice, newSlice)
    for j := 0; j < y; j++ {
      var newInt int
      if j <= count {
        newInt = 0
      } else {
        newInt = 1
      }
      slice[i] = append(slice[i], newInt)
    }
    if i % (d+1) == 0 {
      count++
    }
  }
  shuffle(slice)
  return slice
}

func shuffle(slice [][]int) {
  for i := range slice {
    j := rand.Intn(i + 1)
    slice[i], slice[j] = slice[j], slice[i]
  }
}

// This is used just to write strings to the screen. Used in the "menu".
func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), []rune{}, style)
  }
}

//------------------------------------------------------------------------------
// SORTING ALGORITHM
//------------------------------------------------------------------------------

func quickSort(s tcell.Screen, style tcell.Style, arr [][]int, start, end int) {
  if start >= end {
    return
  }

  index := partition(s, style, arr, start, end)
  quickSort(s, style, arr, start, index - 1)
  quickSort(s, style, arr, index + 1, end)
}

func partition(s tcell.Screen, style tcell.Style, arr [][]int, start, end int) int{

  pivotIndex := start
  pivotValue := arr[end]
  for i := start; i < end; i++ {
    if countLength(arr[i], s) < countLength(pivotValue, s) {
      swap(arr, i, pivotIndex)
      pivotIndex++
    }
  }

  // THIS IS WHERE I WANT TO UPDATE DRAW
  draw(arr, s, style)
  writeToScreen(s, style, 1, 1, "Press 1 For new array")
  writeToScreen(s, style, 1, 2, "Press Q to quit")

  s.Sync()
  time.Sleep(time.Millisecond * 100)

  swap(arr, pivotIndex, end)
  return pivotIndex
}

func swap(arr [][]int, index1, index2 int) {
  temp := arr[index1]
  arr[index1] = arr[index2]
  arr[index2] = temp
}

//------------------------------------------------------------------------------
// SORTING ALGORITHM END
//------------------------------------------------------------------------------


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
