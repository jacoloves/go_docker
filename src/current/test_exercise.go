package main

import (
	"fmt"
)

func main() {
	A := 2.0
	//tmp := 0.0
	//tmp_diff := 100100100.0

	for i := 0.1; i <= 10.0; i += 0.1 {
		A -= (A*A - i) / (2 * A)

		/*if tmp_diff > (A - tmp) {
			tmp_diff = A - tmp
		}

		tmp = A*/

		fmt.Println(A)

	}
}
