package main

import "fmt"

func main()  {
	var nums=[]int{1,2,5,6,3,4,9}
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
}
