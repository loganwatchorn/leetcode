package main

import "fmt"

func main() {
	tests := [][2]string{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
	}

	for _, test := range tests {
		if reverseWords(test[0]) == test[1] {
			fmt.Println("pass")
		} else {
			fmt.Println("fail")
		}
	}
}
