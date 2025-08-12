function strStr(haystack, needle) {
    if (needle === "") return 0;


    const lps = Array(needle.length).fill(0);
    let len = 0, i = 1;
    while (i < needle.length) {
        if (needle[i] === needle[len]) {
            lps[i++] = ++len;
        } else if (len > 0) {
            len = lps[len - 1];
        } else {
            i++;
        }
    }
    let h = 0, n = 0;
    while (h < haystack.length) {
        if (haystack[h] === needle[n]) {
            h++; n++;
            if (n === needle.length) return h - n;
        } else if (n > 0) {
            n = lps[n - 1];
        } else {
            h++;
        }
    }
    return -1;
}
