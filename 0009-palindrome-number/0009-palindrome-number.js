/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    let b=x;
    let a=0;
    while(x>0){
        let i=x%10;
        a=a*10+i;
        x=Math.floor(x/10);
    }
    if(a===b){
        return true;
    }else{
        return false
    }  
};