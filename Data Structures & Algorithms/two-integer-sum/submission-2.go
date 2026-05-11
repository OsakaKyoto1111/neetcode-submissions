func twoSum(nums []int, target int) []int {
    for i,num := range nums{ 
        res:=target-num;
            for index,val:=range nums{ 
                if val==res && i!=index{ 
                    return []int {i,index}
                }
            }

    }
    return []int {}
}
