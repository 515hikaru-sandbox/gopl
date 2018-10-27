package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34
	*/
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Printf("%v\n", ages)
	fmt.Printf("%d\n", ages["alice"])
	// delete(ages, "alice")
	// fmt.Printf("%v\n", ages)
	fmt.Printf("%v\n", ages)
	ages["bob"]++
	fmt.Printf("%v\n", ages)
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	var foo map[string]int
	fmt.Printf("%t\n", foo == nil)
	fmt.Printf("%t\n", len(foo) == 0)
}
