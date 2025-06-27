package main


import ("fmt")
// ? by value
// func changeNum(num int){
	
// 	num =  5
// 	fmt.Println("In changeNum before", num)
// 	// fmt.Println("In changeNum ", num)
// }


// ? by reference

func changeNum(num *int){
	fmt.Println(num)
	*num = 5
	fmt.Println(num)

	fmt.Println("In changeNum", num)




}

func main(){
	num:= 1
	changeNum(&num)
		fmt.Println("memory address of num", &num)
	fmt.Println("After changeNum in main", num)


}