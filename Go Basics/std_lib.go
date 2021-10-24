package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println()

	ages := []int{56, 65, 22, 54, 98, 67, 123, 22, 43, 88, 21}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 98)
	fmt.Println(index)

	names := []string{"Talha", "Abhinav", "Rakesh", "Avinash", "Rajan"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Talha"))

}
