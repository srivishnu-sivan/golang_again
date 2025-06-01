package main

import ("fmt")

//  for  -> only construct  in go for looping
func Loops(){

    // ? while Loop
    i:= 1
    for i<=3{
        fmt.Println(i)
        i= i+1
    }

    for {
        fmt.Println("inifinte loop")
    }

}