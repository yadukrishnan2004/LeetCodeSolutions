func containsNearbyDuplicate(nums []int, k int) bool {
    index:=make(map[int]int)

    for i,num:=range nums{
        if pre,ok:=index[num];ok{
            if i-pre<=k{
                return true
            }
        }

        index[num]=i

    }
    return false
}