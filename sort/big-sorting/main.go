package main
import (
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
    nums := make([]string, n)
    var line string
    
    for ix := 0; ix < n; ix++ {    
        fmt.Scanf("%s", &line)
        nums[ix] = line
    }

    sort.Slice(nums, func(i, k int) bool {
        a, b := nums[i], nums[k]
        // fmt.Println("a", len(a), a, "b", len(b), b)
        if diff := len(a) - len(b); diff != 0 {
            return diff < 0
        }
        
        // fmt.Println(a, len(a), b, len(b))
        for ix := 0; ix < len(a); ix++ {
            // fmt.Println(a[ix], b[ix])
            if a[ix] < b[ix] {
                return true
            }
            if a[ix] > b[ix] {
                return false
            }
        }
        return false
    })
    
    for _, a := range nums {
        fmt.Println(a)
    }
}
