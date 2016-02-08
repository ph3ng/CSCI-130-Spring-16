package main
import "fmt"

func main() {
  //making a slice with composite literal
  x := []int {0,1,2,3}
  //making a slice with make
  y := make([]int, 4)
  //speedier version w/o memory management overhead
  //x := [] int, 0, 4
  y[0] = 5
  y[1] = 4
  y = append(x,5)
  //':' slicing a slice; starting @ 0 but not including the end
  //int type value including variables
  //x = [0:]

  smallest := x[0]
  // '_' char is your throw away var that won't be recorded
  for _, v := range y{
    if v < smallest {
      smallest = v
    }
  fmt.Println(v)
  }
  fmt.Println(y)
  //a slice contains 3 attributes: pointer to an array, size of array, and capacity
  //a slice is an more robust array
}

func a(t []int) {
    panic("Error")
}
