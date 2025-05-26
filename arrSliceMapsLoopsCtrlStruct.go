package main



import ("fmt";
    "time"
)


func ArrSliceMapsLoopsCtrlStruct(){
    // ? Array
    var intArr [3]int32
    intArr[1] = 123
    fmt.Println(intArr[1:3])


    fmt.Println(&intArr[0])
    fmt.Println(&intArr[1])
    fmt.Println(&intArr[2])

    var intArr2 [3]int32 = [3]int32{1,2,3}
    fmt.Println(intArr2)

    intArr3 := [3]int32{2,5,6}
    fmt.Println(intArr3) 


    // ? Slice ===> Slices wraps arrays to give a more general, powerful, and convenient interface to sequence of data
    var intSlice []int32 = []int32{1,2,36,6,2}
    fmt.Println("Slice 1>>>>", intSlice)
    fmt.Printf("The length >> %v and the capacity %v of the Slice 1(intSlice) before appending>>>\n", len(intSlice), cap(intSlice))
    intSlice = append(intSlice,56)
    fmt.Println("Slice 1 after appending >>>>", intSlice)
    // ? to check the capacity of the Slice
    fmt.Println("The length of the Slice 1(intSlice) after appending>>>", len(intSlice))
    fmt.Println("The capacity of the Slice 1(intSlice) after appending>>>", cap(intSlice))
    // eg [1 2 36 6 2 56 * * ], this is how it'll be, we cannot access value which is * in this case, if we try>>> it throws index out of range error index out of range [7] with length 6
    // fmt.Println(intSlice[7])

    // ? we can append multiple values by appending by using spread operator
    var intSlice2 []int32 = []int32{12,56,775}
    fmt.Println("Before using spread operator")
    intSlice = append(intSlice, intSlice2...)
    fmt.Println("After using spread operator")
    fmt.Println("Slice 1>>>>", intSlice)
    // ? The way of creating a slice is by using make function
    intSlice3 := make([]int32, 5,10) // make([]type, length, capacity)
    fmt.Println("Slice 3 created using make function >>>>", intSlice3)


    // ? Map
    var intMap map[string]uint8 = make(map[string]uint8)
    fmt.Println("Map 1 before adding values >>>>", intMap)

    var myMap2 = map[string]uint8{"vishnu":28, "karthi": 29, "sivan":28}
    fmt.Println("Map2 of age>>>", myMap2["sivan"])
    // map alway returns zero value if the key is not present in the map
    fmt.Println("The value what im gonna ask is not in the map, so it returns zero value>>>", myMap2["vishnu1"]) //this is sometimes risk
    // to avoid this we can use comma ok idiom
    var age, ok = myMap2["vishnu"] // change the value to check ["vishnu"]/["vishnu1"]
    if ok{
        fmt.Println("The value of vishnu is >>>", age)
    }else{
        fmt.Println("The key vishnu is not present in the map, so it returns zero value>>>", age)
    }
    // ? deleting a key from the map
    delete(myMap2, "vishnu")
    fmt.Println("Map 2 after deleting vishnu >>>>", myMap2)

    // ? Looping through the map
    people := map[string]int{
        "Alice":    25,
        "Bob":      30,
        "Charlie":  28,
        "Diana":    22,
        "Ethan":    35,
        "Fiona":    27,
        "George":   32,
        "Hannah":   24,
        "Ian":      29,
        "Jessica":  26,
    }

    // ? the order is not preserver in the map, so it will print in random order
    fmt.Println("Looping through the map >>>>",people)
    for name :=range people{
        fmt.Println("The name is >>>>", name)

    }

    for name , age :=range people{
        fmt.Printf("The name is %v and the age is %v\n", name ,age)
    }

    // ? there is no while loop in go, but we can use for loop as while loop
    var i int = 0
    for i<=10{
        fmt.Println("The value of i is >>>>", i)
        i++
    } //  this one way of doing while loop


    for {
        if i >=10{
            break
        }
        fmt.Println("The value of i is >>>>", i)
        i++
    }

    // lets check the time consumption for the slice 
    var n int  = 1000000
    var testSlice  = []int{}
    var testSlice2  = make([]int, 0, n) // this is how we can create a slice with capacity n
    fmt.Printf(" Total time without preallocation : %v\n", timeLoop(testSlice, n))
    fmt.Printf(" Total time with preallocation : %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration{
    var t0 = time.Now()
    for len(slice)<n{
        slice = append(slice,1)
    }
    return time.Since(t0)
}