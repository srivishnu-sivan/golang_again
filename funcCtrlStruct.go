package main

import ("fmt"
        "errors"
)

func FuncCrtlStrct(data myDataStruct) {
    printME(data.dataToPrint)

    result, err := intDivision(data.numerator, data.denominator)
    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Result of division:", result)
    }
}

func printME(printValue string){
    fmt.Println("printValue from params", printValue)
}

func intDivision(numerator int, denominator int) (int, error) {
    
    if denominator==0{
        // error.New("Cannot Divide by the value ZERO")
        return 0, errors.New("Cannot Divide by the value ZERO")
    }
    var result int = numerator/denominator
    fmt.Printf("\nThe result of the integer division is %v and %v is %v \n",numerator,denominator, result)
    return result,nil
}

