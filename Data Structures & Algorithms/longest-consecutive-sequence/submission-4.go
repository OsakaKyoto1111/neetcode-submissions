func longestConsecutive(nums []int) int {
	m:=map[int]bool{}
	count:=0
	for _,val:=range nums{
		m[val] = true;
	}

		for _,val :=range nums{ 
			if !m[val-1]{ 
				lens:=1
				current:=val
				for m[current+1]{ 
					lens++
					current++
				}
				if lens > count{ 
					count = lens 
				}
			}
		}
		return count
}
