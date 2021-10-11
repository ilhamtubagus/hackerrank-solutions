package solution
// https://www.hackerrank.com/challenges/kangaroo/problem
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    if v2<v1 {
        xs := x2-x1
        vs := v2-v1
        if xs%vs == 0 {
            return "YES"
        }else{
            return "NO"
        }
    }else{
		//the first kangooro will never be able to catch up
        return "NO"
    }
}