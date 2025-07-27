package main

import "fmt"


func addTen(num *int) {
    *num = *num + 10 
}

func main() {
    value := 5
    fmt.Println("调用前:", value) 
    
    addTen(&value)     
    
    fmt.Println("调用后:", value) 
}
