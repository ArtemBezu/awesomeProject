package main

import (
	"fmt"
)

func simplifiedFractions(n int) []string {
	var res []string

	for denominator := 1; denominator <= n; denominator++ {
	NumLoop:
		for numerator := 1; numerator < denominator; numerator++ {
			if numerator > 1 && denominator%numerator == 0 {
				continue
			}
			for i := 2; i <= denominator/2; i++ {
				if numerator%i == 0 && denominator%i == 0 {
					continue NumLoop
				}
			}
			current := fmt.Sprintf("%d/%d", numerator, denominator)
			res = append(res, current)
		}
	}
	return res
}
