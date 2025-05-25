package main

import ("fmt"; "unicode/utf8")

func Variables(){
    fmt.Println("Hey, It's Golang again")
    // ? Data types

    var intNum int = 3278962536
    fmt.Println("intNum", intNum)

    var floatNum float32 = 12.36
    fmt.Println("floatNum", floatNum)


    var floatNum32 float32 = 12.253
    var intNum32 int32 = 2
    // ? it's not directly possible to add an integer and float in golang just like javascript
    var result float32 =  floatNum32 + float32(intNum32)
    fmt.Println("result>>>>", result)


    var intNum1 int =5
    var intNum2 int =2
    fmt.Println("/",intNum1/intNum2)
    fmt.Println("%",intNum1%intNum2)

    // ? string
    // var myString string = "Hello world"
    // var myString string = "Hello" + " world"
    var myString string = "ÿ"
    fmt.Println("myString>>>", myString)
    // ? since go utf8 encoding characters outside the vanilla ascii characters are stored with more than one byte
    fmt.Println("length of myString>>>>>>.", len(myString))
    // ? so to get the actual length of the string import a package call unicode/utf8 and use it's inbuilt function(Rune are another datatypes in go)
    fmt.Println("lenght of the myString>>>>>", utf8.RuneCountInString(myString))
    var myRune rune = 'a'
    fmt.Println(myRune)

    var myBoolean bool =  false
    fmt.Println("myBoolean>>>",myBoolean)

    myVar := "Some Random Text!!!"
    var1, var2 := 1,2
    fmt.Println(myVar)
    fmt.Println(var1, var2)


    const myConst string = "const value"
    // myConst =  "another const value"   // this is cannot happen just like  javascript
    fmt.Println("myConst>>",myConst)
}