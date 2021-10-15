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
 
 func reduce(f func(v1, v2 int32) int32, a []int32) int32{
	 var result int32
	  if len(a) == 1 {
			 x := a[0]
			 y := int32(0)
			 return f(x, y)
		 }
	 j:=1
	 for i:= 0; i < len(a) - 1; i+=2{
		 result = f(a[i], a[j])
		 j++
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