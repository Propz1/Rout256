package main

import (
	"bufio"
	"fmt"
	"os"

	"sync"
)

var (
	set int
	pr  int
	mu  sync.RWMutex
	numDiscount int = 2
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	AmountToPaid(reader)
}

func AmountToPaid(reader *bufio.Reader) {

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var countSets uint16
	fmt.Fscan(reader, &countSets)

	var totalPrice = make([]int, 0)

	var i uint16 = 1
	max := 2*countSets + 1

	for {

		if i == max {
			break
		}

		if i%2 == 0 { //even

			price := make(map[int]int, set)

			for n := 0; n < set; n++ {

				fmt.Fscan(reader, &pr)

				mu.Lock()
				price[pr]++
				mu.Unlock()

			}

			var tp int

			for p, count := range price {
				tp = tp + (numDiscount*(count/3)+count%3)*p
			}

			mu.Lock()
			totalPrice = append(totalPrice, tp)
			mu.Unlock()
		}

		if i%2 != 0 { //odd
			fmt.Fscan(reader, &set)
		}

		i++
	}

	for _, tp := range totalPrice {
		fmt.Fprintln(output, tp)
	}

}
