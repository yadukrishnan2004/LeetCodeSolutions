func lengthOfLongestSubstring(s string) int {
        k:=make(map[byte]int)
        b,l:=0,0
        for i:=0;i<len(s);i++{
            ch:=s[i]
            if idx,ok:=k[ch];ok&&idx>=l{
                l=idx+1
            }
            k[ch]=i
            if i-l+1>b{
                b=i-l+1
            }
        }
        return b
}