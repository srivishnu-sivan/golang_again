package main


import ("fmt"
"time"

    )


func Switch(){
   
	i := 8

	switch i {
	case 1:
		fmt.Println("i is !")
	case 2:
		fmt.Println("i is 2")

	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("Other")
	}

    switch time.Now().Weekday(){
        case time.Saturday, time.Sunday:
            fmt.Println("It's weekend !!!")
        case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
            fmt.Println("It's a weekday")
        default:
            fmt.Println("It's a day of the week")
    }

}