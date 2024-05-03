package main

import "fmt"


func addTen(a int) int {
  return a + 10;
}

func multiplyTwo(a int) int {
  return a * 2;
}

func main(){

  num := 10;

  fmt.Println(num); //println(num);

  var pipeline = [...](func(int)int){addTen, multiplyTwo, multiplyTwo};

  for index, el := range pipeline{
    fmt.Println("Executing ", index)
    num = el(num);
  }

  fmt.Println(num);
}
