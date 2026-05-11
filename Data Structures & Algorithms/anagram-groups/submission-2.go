func groupAnagrams(strs []string) [][]string {
	res:= [][]string{} 
	groups:=map[[26]int][]string{}
	for _,val:=range strs { 
		counter:=[26]int{}
			for _,ch:=range val{ 
				idx:=int(ch-'a')
				counter[idx]++
			}
		groups[counter] = append(groups[counter],val)

	}
	for _,group:=range groups{ 
		res = append(res,group)

	}
	return res
}
