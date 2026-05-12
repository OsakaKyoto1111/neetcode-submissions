func maxArea(heights []int) int {
	min:=0
	sum:=0
	maxArea:=0

		left :=0
		right:=len(heights)-1
			for left<right{
				if heights[left]<=heights[right]{
					min = heights[left]
					sum = min*(right-left)
					left++
				}else{
					min = heights[right]
					sum = min*(right-left)
					right--
				}
				current:=sum


						
				if current>maxArea{
				maxArea = current
			}
			}
	

				
	
	return maxArea
}
