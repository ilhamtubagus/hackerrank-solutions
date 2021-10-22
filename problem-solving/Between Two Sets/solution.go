package solution

func getGCD(x, y int32) int32 {
	for x % y > 0 { 
		remainder := x % y
		x = y
		y = remainder
	}
	return y
 }
 func lcm(x, y int32) int32{
	 return (x*y)/getGCD(x,y)
 }
func reduce(a []int32, f func(v1, v2 int32) int32) int32 {
	result := a[0]
	var ins [2]int32
	for i := 0; i < len(a); i++ {
		ins[0] = result
		ins[1] = a[i]
		result = f(ins[0], ins[1])
	}
	return result
}

 func GetTotalX(a []int32, b []int32) int32 {
	 // Write your code here
	 
	 f1 := reduce(lcm, a)
	 f2 := reduce(getGCD, b)
	 var result int32
	 for i:= f1; i <= f2; i+=f1 {
		 if f2%i == 0 {
			 result++
		 }
	 }
	 return result
 }
