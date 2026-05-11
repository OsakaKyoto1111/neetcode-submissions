func isAnagram(s string, t string) bool {
	var a [26]int
	var b [26]int

	for _,val:=range s {
		idx := int(val - 'a')
		a[idx]++
	}
		for _,valB:=range t {
		idxB := int(valB - 'a')
		b[idxB]++
	}
	if a == b {
		return true
	}else{
		return false
	}
}
