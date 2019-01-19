package main

import (
	"fmt"
)

type programlang struct {
	name     string
	category string
	rank     int
}

type coder struct {
	personname string
	programlang
}

func main() {
	pg1 := programlang{
		name:     "C++",
		category: "Static",
		rank:     2,
	}

	pg2 := programlang{
		name:     "Javascript",
		category: "Dynamic",
		rank:     1,
	}
	fmt.Println(pg1)
	fmt.Println(pg1.name, " - ", pg1.category, " - ", pg1.rank)

	fmt.Println(pg2)
	fmt.Println(pg2.name, " - ", pg2.category, " - ", pg2.rank)

	coder1 := coder{
		personname:  "Dam",
		programlang: pg1,
	}
	fmt.Println(coder1)
	fmt.Println(coder1.personname, " - ", coder1.programlang)

	coder2 := struct {
		personname string
		role       string
		programlang
	}{
		personname:  "Dam",
		role:        "Senior Software Engineer",
		programlang: pg2,
	}
	fmt.Println(coder2)
	fmt.Println(coder2.personname, " - ", coder2.programlang)
}
