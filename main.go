package main

import (
  "fmt" 
  "strconv"
)
//Address structure
type Address struct{
  _number int
  _street string
  _zip int 
  _city string
  _state string
  _country string
}

func main() {
  printValues()
  printUndefinedValues()
  var address Address = createStructures(123, "fake street", 32808, "Orlando", "FL", "USA" )
  printAddress(address)
  makeArray()
}
//Basic print of hard coded values
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
//Basic print of undefined values
func printUndefinedValues(){
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
//Function to create and fill the address structure
func createStructures(number int, street string, zip int, city string, state string, country string)Address{
  retVal  := Address{number, street, zip, city, state,country}

  return retVal
}
//Function to format and print the address structure
func printAddress(address Address){
  fmt.Println(address._number, " ", address._street, " ", address._city, ", ", address._state, " ", address._country )
}
//Function to make a generic array
//This function is dumb, but it serves the purpose of
//creating an array, printing said array, and casting int to string if an empty value is found
func makeArray(){
  const length int = 10
  var arr[length]string
  arr[0] = "0"
  arr[1] = "1"
  arr[2] = "2"
  arr[3] = "3"
  arr[4] = "4"

  for i:=0; i < length; i++{
    if len(arr[i]) != 0{
      fmt.Println(arr[i])
    }else{
      arr[i] = strconv.Itoa(i)
      fmt.Println(arr[i])
    }
  }
}