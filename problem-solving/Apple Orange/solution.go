package solution
// https://www.hackerrank.com/challenges/apple-and-orange/problem
import (
    "fmt"
)

func inBetween(v, min, max int32) bool {
    if (v <= max) && (v >= min){
        return true
    }else{
        return false
    }
}
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    // Write your code here
    var aLand []int32
    var oLand []int32
    
    for _, appl := range apples {
        applLand := appl+a
        if inBetween(applLand, s, t) {
        aLand = append(aLand, applLand)
        }
    }
    for _, orange := range oranges{
        orangeLand := orange+b
        if inBetween(orangeLand, s, t){
                    oLand = append(oLand, orange+b)

        }
    }
    fmt.Println(len(aLand))
    fmt.Println(len(oLand))
}