//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Quick Sort Demo 2022
//
// Author: Tyler(UnclassedPenguin)
//    URL: https://unclassed.ca
// GitHub: https://github.com/UnclassedPenguin
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)


// Main starting point. Tells the user how to run the program, and sets up the
// initial array, after the user has selected.
func menu(s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  var slice [][]int
  divide := int(x/y)+1
  var sorted bool

  strings := []string{ "Unclassed Penguin Quick Sort",
                       "Press 1 to start from random seed",
                       "Press 2 to start from shuffled \"ordered\" seed",
                       "Esc, Ctrl-C, or q to quit",
                     }

  // Write strings to screen.
  for i, str := range strings {
    writeToScreen(s,style,((x/2)-(len(str)/2)),y/3+(i*2),str)
  }

  // Keyboard handling. Keys to quit (Esc, Ctrl-c, q)
  // and the key to start the game (1,2,3)
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
          sorted = false
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
          sorted = false
          slice = createOrderedSlice(s, divide)
          draw(slice, s, style)
          writeToScreen(s, style, 1, 1, "Press 1 For new random array")
          writeToScreen(s, style, 1, 2, "Press 2 For new \"ordered\" array")
          writeToScreen(s, style, 1, 3, "Press 3 to sort")
          writeToScreen(s, style, 1, 4, fmt.Sprintf("X: %v", x))
          writeToScreen(s, style, 1, 5, fmt.Sprintf("Y: %v", y))
          writeToScreen(s, style, 1, 6, fmt.Sprintf("d: %v", divide))
          s.Sync()

        case '3':
          if !sorted {
            quickSort(s, style, slice, 0, len(slice)-1)
            sorted = true
          } else {
            writeToScreen(s, style, x/2-7, 1, "Already Sorted!")
            s.Sync()
          }
        }
      }
    }
  }
}

// Returns a count of how many 1's are in any given array.
// Effectively, gives you the size of each column, so that it
// can be sorted.
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


// Draws an array of arrays to the screen.
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

// Function used to create random slice for the terminal window. 
// For each column it picks a random number between 0 and the height
// of the column.
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

// This function creates an "ordered" slice. A slice that increases
// by a known amount based on terminal width/height.
// "d" is (x/y)+1 so, (width/height)+1
// "d" just describes how many columns will have the same height before
// The height is increased by 1. 
func createOrderedSlice(s tcell.Screen, d int) [][]int {
  x, y := s.Size()

  var slice [][]int
  count := 1

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
    if i % (d) == 0 {
      count++
    }
  }
  shuffle(slice)
  return slice
}

// Shuffles the slice of slices to a random order. 
// Shuffles the columns.
func shuffle(slice [][]int) {
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(slice), func(i, j int){
    slice[i], slice[j] = slice[j], slice[i]
  })
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

// Starts the recursive quick sorting. Starts over the whole array, then splits
// it and continues to split until it is done.
func quickSort(s tcell.Screen, style tcell.Style, arr [][]int, start, end int) {
  if start >= end {
    return
  }

  index := partition(s, style, arr, start, end)
  quickSort(s, style, arr, start, index - 1)
  quickSort(s, style, arr, index + 1, end)
}

// This is where the real sorting happens. It takes the last element(the pivotValue) 
// of its given chunk, and iterates over the entire array, at each step checking 
// wether the element is less than the pivotValue. If it is, then it is swapped
// with the pivotIndex (which starts at the very beginning of the section), and the
// pivotIndex is increased by 1. after doing this over every element of the section,
// the final step is to swap the last element(The pivotValue) with whatever element
// is in the pivotIndex. After this, all elements to the left are lower than the
// pivotIndex, and all elements to the right of the pivot index are higher than the 
// pivotIndex, although are still completely unsorted, so this is where it 
// recursively calls itself and sorts the divided sections again.
func partition(s tcell.Screen, style tcell.Style, arr [][]int, start, end int) int{
  pivotIndex := start
  pivotValue := arr[end]

  for i := start; i < end; i++ {
    if countLength(arr[i], s) < countLength(pivotValue, s) {
      swap(arr, i, pivotIndex)
      pivotIndex++
    }
  }

  swap(arr, pivotIndex, end)

  draw(arr, s, style)
  writeToScreen(s, style, 1, 1, "Press 1 For new random array")
  writeToScreen(s, style, 1, 2, "Press 2 For new \"ordered\" array")
  writeToScreen(s, style, 1, 3, "Press Q to quit")

  s.Sync()
  time.Sleep(time.Millisecond * 100)

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


// Main function. Initializes screen and passes to the "menu" func. 
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
