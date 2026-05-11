func isAnagram(s string, t string) bool {
    charsA:=[]rune(s)
    charsB:=[]rune(t)
    
    for i:=0 ; i<len(charsA);i++{
        for j:=i+1 ; j<len(charsA);j++{
            if charsA[i]>charsA[j]{
                charsA[i],charsA[j] = charsA[j],charsA[i]
            }
        }
    }
     for k:=0 ; k<len(charsB);k++{
        for l:=k+1 ; l<len(charsB);l++{
            if charsB[k]>charsB[l]{
                charsB[k],charsB[l] = charsB[l],charsB[k]
            }
        }
    }
    resultA := string(charsA);
        resultB := string(charsB);
        if resultA ==resultB { 
            return true;
        }else{ 
            return false
        }

}
