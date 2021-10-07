package solution

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SuperStack(operations []string) {
	stack := []int{0}
	for _, v := range operations {
		o := strings.Split(v, " ")
		if o[0] == "pop" {
			//pop from the stack
			stack = stack[:len(stack)-1]
		} else if o[0] == "push" {
			val, err := strconv.Atoi(o[1])
			if err != nil {
				log.Fatal("value must be an integer")
			}
			//push onto stack
			stack = append(stack, val)
		} else if o[0] == "inc" {
			//add v to each bottom i element
			i, err := strconv.Atoi(o[1])
			if err != nil {
				log.Fatal("value i must be an integer")
			}
			v, err := strconv.Atoi(o[2])
			if err != nil {
				log.Fatal("value v must be an integer")
			}
			for x := 0; x < i && x < len(stack); x++ {
				stack[x] += v
			}

		}
		if stack[len(stack)-1] == 0 {
			fmt.Println("EMPTY")
		} else {
			fmt.Println(stack[len(stack)-1])
		}
	}
}
