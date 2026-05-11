func isPalindrome(s string) bool {	
	final:=[]rune{}
	for _,ch:=range s { 
		if ch >='a' && ch <='z'{
			final = append(final,ch)
		}else if ch >='A' && ch <='Z'{
			ch = ch+32
			final = append(final,ch)
		}else if ch >='0' && ch <='9'{
			final = append(final,ch)
		}else{
			continue
		}
	}
	left:=0
	right:=len(final)-1 
	for left<right{
		if final[left] != final[right]{
			return false

		}
			left++
			right--
	}
	 return true



}



