package main

import (
	"fmt"
	"regexp"
)

func main() {
	strs := []string{
		"{\"text\":\"jenkins --help\"}",
		"{\"text\":\"@_user_1 --help\"}",
		"{\"text\":\"--help\"}",
		"{\"text\":\"      --help\"}",
	}
	re := regexp.MustCompile(`^{"text":"(@_user_1 )?(.*?)"}$`)
	for _, str := range strs {
		match := re.FindStringSubmatch(str)
		fmt.Print(len(match))
		fmt.Print("  ")
		fmt.Println(match)
		fmt.Println(match[0])
		fmt.Println(match[1] == "")
		fmt.Println(match[2])
	}

}
