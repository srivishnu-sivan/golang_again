package main


import (
    "fmt"
    // "math/rand"
    "time"
    "sync"
)

var dbData = []string{"apple", "banana", "cherry", "date"}
var resultSlice []string
var mu = sync.Mutex{}

func Goroutines(){
    fmt.Println("func check")
    var wg sync.WaitGroup

    t0 := time.Now()
    for i:=0; i<len(dbData); i++{
        // dbCall(i)// if you use go keyword in front of the function
        wg.Add(1)
       go func(i int) {
         defer wg.Done()   // ✅ Inside goroutine — gets called when goroutine ends
        dbCall(i)
    }(i)
    }
    //  time.Sleep(1 * time.Second)
    // wg.Wait()
    fmt.Printf("\nTotal execution time : %v\n", time.Since(t0))
}


func dbCall(i int ){
    var delay float32 = 2000
    // mu.Lock() // Lock the mutex to ensure exclusive access to the shared resource
    time.Sleep(time.Duration(delay)*time.Millisecond)
    fmt.Println(" the Result from the database is : ", dbData[i])
    resultSlice = append(resultSlice, dbData[i])
    // mu.Unlock() // Unlock the mutex to allow other goroutines to access the shared resource
    fmt.Println(" the Result from the database is : ", resultSlice)
    
}