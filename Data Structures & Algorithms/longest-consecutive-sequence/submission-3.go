func longestConsecutive(nums []int) int {
	m:=map[int]bool{}
	count := 0
	for _ ,val :=range nums{
		m[val] = true;
	}

		for _,val :=range nums{ 
		 if !m[val-1]{
			cur := val
			lens := 1
			for m[cur+1]{
				cur++
				lens++
			}
			if lens > count { 
				count = lens
			}
		 }

		}
		return count

}
