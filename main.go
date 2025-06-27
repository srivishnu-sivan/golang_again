package main

import ("fmt")


type myDataStruct struct {
    dataToPrint string
    numerator int
    denominator int
} 



func main(){
    // Variables()
    // paramVal := myDataStruct{dataToPrint : "Hello, I'm actually coming from different file", numerator : 25, denominator : 5}
    // FuncCrtlStrct(paramVal)
    // ArrSliceMapsLoopsCtrlStruct()
    // StructInterface()
    // Goroutines()
    // AddNum()
    // Loops()
    // Switch()



    // Example 1
	nums1 := []int{2, 7, 11, 15,6,3}
	target1 := 9
	// fmt.Println(LeetCode(nums1, target1)) // Output: [0 1]

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(LeetCode(nums2, target2)) // Output: [1 2]

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(LeetCode(nums3, target3)) // Output: [0 1]

}