package main

func reverseString(s []byte)  {
    n := len(s) - 1;
    arr := make([]byte, n + 1);
    
    for i := 0; i <= n; i++ {
        arr[i] = s[n - i];
    } 
    
    copy(s, arr);
}