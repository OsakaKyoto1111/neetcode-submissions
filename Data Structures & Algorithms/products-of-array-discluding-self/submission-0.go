func productExceptSelf(nums []int) []int {
	answer:=make([]int,len(nums));
	for i:=0;i<len(nums);i++{
		answer[i]=1
		for j:=0;j<len(nums);j++{
			if i!=j{
				answer[i]*=nums[j]
			}
		}
	}
	return answer
}
