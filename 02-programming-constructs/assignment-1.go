/* Print all the prime numbers between 3 & 100 */

package main

import "fmt"

func main() {
	for i := 3; i <= 100; i++ {
		isPrime := true
		for j := 2; j <= (i / 2); j++ {

			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}
