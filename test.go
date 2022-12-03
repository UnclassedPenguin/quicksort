package main

import (
  "fmt"
  "time"
  "math/rand"
)

func quickSort(arr []int, start, end int) {
  if start >= end {
    return
  }

  index := partition(arr, start, end)
  quickSort(arr, start, index - 1)
  quickSort(arr, index + 1, end)
}

func partition(arr []int, start, end int) int{

  pivotIndex := start
  pivotValue := arr[end]
  for i := start; i < end; i++ {
    if arr[i] < pivotValue {
      swap(arr, i, pivotIndex)
      pivotIndex++
    }
  }
  swap(arr, pivotIndex, end)
  return pivotIndex
}

func swap(arr []int, index1, index2 int) {
  temp := arr[index1]
  arr[index1] = arr[index2]
  arr[index2] = temp
}

func main() {

  rand.Seed(time.Now().UnixNano())
  var arr []int
  length := 40

  for i := 0; i < length; i++ {
    arr = append(arr, rand.Intn(length))
  }

  fmt.Println("arr:", arr)

  quickSort(arr, 0, len(arr)-1)

  fmt.Println("sorted:", arr)

}
