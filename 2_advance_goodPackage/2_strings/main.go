package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings Test Start...")

	//var b strings.Builder
	//b.Grow(17)

	var names []string
	names = append(names, "James", "Wade", "Paul", "Anthony")
	//name := strings.Join(names, "-")
	//fmt.Printf("[%s]\n", name)

	n := 1 * len(names) - 1
	for i := 0; i < len(names); i++ {
		n += len(names[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(names[0])
	for _, v := range names[1:] {
		b.WriteString("-")
		b.WriteString(v)
	}
	fmt.Printf("[%s]\n", b.String())
}