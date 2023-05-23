package main

func minStartValue(nums []int) int {
    for i := 1; ; i++ {
        s := i;
        flag := true;
        
        for _, e := range nums {
            s += e
            if s < 1 {
                flag = false
                break;
            }
        } 
        
        if (flag) {
            return i;
        }
    }
}

