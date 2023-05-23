package main

func selfDividingNumbers(left int, right int) []int {
    var arr []int;
    
    for ; left <= right; left++ {
        if (check(left)) {
            arr = append(arr, left);
        }
    }
    
    return arr;
}

func check(n int) bool {
    for t := n; t != 0; t /= 10 {
        d := t % 10;
        if (d == 0 || n % d != 0) {
            return false;
        }
    }
    
    return true;
}