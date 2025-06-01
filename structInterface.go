package main

import ("fmt")

type gasEngine struct{
    mpl uint
    tankCapacity uint
}

func StructInterface(){
    var myEngine gasEngine = gasEngine{mpl:25, tankCapacity:25}
    fmt.Println("first struct>>>", myEngine)
}