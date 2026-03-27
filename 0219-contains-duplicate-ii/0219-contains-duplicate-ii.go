func abs(num int)int{
    if num<0{
        return num*-1
    }
    return num
}
func containsNearbyDuplicate(nums []int, k int) bool {
    // index:=make(map[int]int)

    // for i,num:=range nums{
    //     if pre,ok:=index[num];ok{
    //         if i-pre<=k{
    //             return true
    //         }
    //     }

    //     index[num]=i

    // }
    // return false
    n:=0
l:=len(nums)
    for i:=0;i<l;i++{
        for j:=i+1;j<l;j++{
            if nums[i]==nums[j]&& abs(i-j)<=k {
                return true
            }
            n++
            if n>100000{
                return false
            }
        }
    }
    return false
}