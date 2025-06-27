package main

import (
	"fmt"
	"sync"
	"time"
)



func task(id int, wg *sync.WaitGroup){

	defer wg.Done()
	fmt.Println("worker >>>>", id)
}


func main(){
	fmt.Println(time.Stamp)
	var wg sync.WaitGroup

	for i:=0; i<=4; i++{
		wg.Add(1)
	 go task(i, &wg)
	}

	wg.Wait()

}