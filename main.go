package main

import "fmt"

func main(){
fmt.Println("1 + 1 = ",1.0+1.0)
fmt.Println("Hello World")
fmt.Println(len("Hello World"))
fmt.Println("Hello World"[1])
fmt.Println("Hello " + "World")

fmt.Println("Hello" == "World")
fmt.Println("Hello" == "Hello")

var a int = 22
const b int =22
fmt.Println(a)
fmt.Println(b)
}