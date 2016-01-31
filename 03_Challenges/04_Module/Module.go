package main

import "fmt"

func main (){
      var first_num, second_num int
      fmt.Println("Please enter a number, followed by a numbered larger than the last number.")
      fmt.Scan(&first_num, &second_num)

      fmt.Println(first_num, "%", second_num, "=", first_num%second_num)
}
