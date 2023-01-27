package main

import "fmt"

func main() {

	m := map[string]int{
		"ABC":    29,
		"ABCD":   34,
		"ABCDE":  12,
		"ABCDEF": 78,
		"XYZ":    34,
	}

	for k, v := range m {
		if v == 34 {
			m[k] = v + 1
		}
		fmt.Printf("Key is %s  and  Value is %v \n", k, m[k])
	}
}
