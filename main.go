package main

import ("fmt")

func main() {
	fmt.Println("Hello, World!")
  printValues()
  printUndefinedValues()
}

func printValues(){
  var val string = "Test1"
  var val1 string  = "Test2"
  val2 := "Test3"
  val3 := 123
  fmt.Println(val)
  fmt.Println(val1)
  fmt.Println(val2)
  fmt.Println(val3)
}

func printUndefinedValues(){
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}