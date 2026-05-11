func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	res := [][2]int{}
	
	for _, val := range nums {
		freq[val]++
	}
	for key,val :=range freq{ 
		res = append(res, [2]int{key, val})
	}
		for i:=0;i<len(res);i++{
			for j:=i;j<len(res);j++{
				if res[i][1] > res[j][1] {
					res[i], res[j] = res[j], res[i]
				}}
		}
		
	answer:=[]int{}
	for a:=0;a<k;a++{ 
		answer=append(answer,res[len(res)-1-a][0])
	}
	return answer
	}
		

 
