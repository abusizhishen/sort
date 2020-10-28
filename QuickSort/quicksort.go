package main

func quickSort(nums []int, begin,end int)  {
	if begin>=end{
		return
	}

	pivot := partition(nums,begin,end)
	quickSort(nums,begin,pivot-1)
	quickSort(nums,pivot+1,end)
}

func partition(nums []int, begin,end int)int  {
	pivot,count := end,begin

	for i:=begin;i<end;i++{
		if nums[i]<nums[pivot]{
			nums[i],nums[count] = nums[count],nums[i]
			count++
		}
	}

	nums[pivot],nums[count] = nums[count],nums[pivot]
	return count
}