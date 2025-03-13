package main

import (
	"fmt"
	"regexp"
	"strings"
)

func determineIndex(y int) int {
	if y > 25 {
		return determineIndex(y - 26)
	}

	return y
}

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */
func caesarCipher(s string, k int32) (result string) {
	const lib = "abcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder
	// Write your code here
	for _, val := range s {
		stringVal := string(val)
		pattern := `^[A-Za-z]+$`
		re := regexp.MustCompile(pattern)

		if re.MatchString(stringVal) {
			index := strings.Index(lib, stringVal)
			isConvertedToLower := false
			if index == -1 {
				// not found, convert to lower case and save flag to indicate that need to revert back to upper case
				index = strings.Index(lib, strings.ToLower(stringVal))
				isConvertedToLower = true
			}

			index += int(k)
			if index > 25 {
				index = int(determineIndex(index))
			}

			stringVal = string(lib[index])
			if isConvertedToLower {
				stringVal = strings.ToUpper(stringVal)
			}
			builder.WriteString(stringVal)
		} else {
			builder.WriteString(stringVal)
		}
	}

	return builder.String()
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
 * Complete the 'palindromeIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func palindromeIndex(s string) int32 {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			// Check if removing left character makes palindrome
			if isPalindrome(s, left+1, right) {
				return int32(left)
			}
			// Check if removing right character makes palindrome
			if isPalindrome(s, left, right-1) {
				return int32(right)
			}
			return -1 // If neither makes it palindrome
		}
		left++
		right--
	}

	return -1
}

func main() {
	//fmt.Println(caesarCipher("www.abc.xy", 87))
	fmt.Println("Removed", palindromeIndex("aabaaa"))
	//fmt.Println(strings.Index("aaab", "b"))
}
