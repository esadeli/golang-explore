// Hands-on exercise #2
// Print every rune code point of the uppercase alphabet three times. Your output should look like this:
// 65
//     U+0041 'A'
//     U+0041 'A'
// U+0041 'A'
// 66
//     U+0042 'B'
//     U+0042 'B'
//     U+0042 'B'
//  … through the rest of the alphabet characters
package main

import (
	"fmt"
)

func main4() {
	for i := 30; i <= 40; i++ {
		if i%2 != 0 && i%3 == 0 {
			fmt.Println(`This is i `, i)
			for j := 0; j < 3; j++ {
				fmt.Printf("Gave me the real number %#U\n", i)
			}
		}
	}
}
