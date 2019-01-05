package main

import (
	"fmt"
)

func main() {
	// normal loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// nested loops
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			fmt.Printf("outer loop %d and inner loop %d\n", j, k)
		}
	}

	// while like for - loops
	l := 2
	for l < 6 {
		fmt.Println("L ke ----", l)
		l++
	}
	fmt.Println("DONE")

	// while like for - loops + breaks
	k := 4
	for {
		if k >= 8 {
			break
		}
		fmt.Println("K ke ----", k)
		k++
	}
	fmt.Println("DONE TWICE")

	// while like for - loops + breaks + continue
	m := 3
	for {
		if m > 8 {
			break
		}
		m++
		if m%2 == 0 {
			fmt.Println(m % 2)
			fmt.Println(m / 2)
			continue
		}
		fmt.Println("M ke ----", m)
	}
	fmt.Println("DONE THREE TIMES")

	// print ascii
	for n := 30; n < 38; n++ {
		fmt.Printf("%v -- %#U -- %#x\n", n, n, n)
	}

	// if else - nested
	p := 2
	for {
		if p > 20 {
			break
		} else if p == 2 {
			fmt.Printf("%d is a Prime number\n", p)
		} else if p%2 == 0 {
			fmt.Printf("%d is not a Prime number\n", p)
		} else {
			fmt.Printf("%d is a probably number\n", p)
		}
		p++
	}
	fmt.Println("Finish")
}
