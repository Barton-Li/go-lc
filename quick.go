package main
func quickSort(nums []int)[]int{
	if len(nums)<=1{
		return nums
	}
	pivot:=nums[0]
	var left,right []int
	for _,num:=range nums{
		if num<pivot{
			left=append(left, num)
		}else{
			right = append(right, num)
		}
	}
	return append(append(quickSort(left),pivot),quickSort(right)...)
}