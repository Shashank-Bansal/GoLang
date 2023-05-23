package main

func largestAltitude(gain []int) int {
    var max, alt, n int = 0, 0, len(gain)
    
    for i := 0; i < n; i++ {
        alt += gain[i];
        if (max < alt) {
            max = alt
        }
    }
    
    return max;
}