package main

import ("fmt")

func AddNum(){
    fmt.Println("AddNum function called")
    var Arr []int = []int{34, 91, 13, 76, 28, 42, 85, 19, 67, 51, 23, 98, 11, 46, 72}
    var sumArr int = 0
    var evenArr []int = []int{}
    for i := range Arr{
        if Arr[i] % 2 == 0{
            sumArr +=Arr[i]
            evenArr = append(evenArr, Arr[i])
        }
    }
    fmt.Println("evenArr, evenArr", evenArr)
    fmt.Println("sum of Arr", sumArr)

}


