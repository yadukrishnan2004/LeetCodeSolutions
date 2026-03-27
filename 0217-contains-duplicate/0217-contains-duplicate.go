func containsDuplicate(nums []int) bool {
    m:=make(map[int]bool)
    for _,i:=range nums{

        if m[i]==true{
            return true
        }
        m[i]=true
    }
    return false
}