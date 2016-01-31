package main

import "fmt"

func main() {
      fmt.Println("Please input your full name.")
      var fname, lname string
      fmt.Scan(&fname, &lname)
      fmt.Println(fname, lname)
}
