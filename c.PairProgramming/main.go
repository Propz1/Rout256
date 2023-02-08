package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	set uint8
	n   uint8
	lev uint8
	x   uint8
	y   uint8
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	PairProgramming(reader)

}

func PairProgramming(reader *bufio.Reader) {

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var countSets uint8
	fmt.Fscan(reader, &countSets)

	result := make([]uint8, 0, set)

	var i uint8 = 1
	max := 2*countSets + 1

	for {

		if i == max {
			break
		}

		if i%2 == 0 { //even

			levels := make([]uint8, 0, set)

			for n = 0; n < set; n++ {

				fmt.Fscan(reader, &lev)

				levels = append(levels, lev)
			}

			temp := make(map[uint8]struct{}, set)

			for x = 0; x < set; x++ {

				differ := make(map[uint8]uint8, set)

				if x > 0 {
					if _, ok := temp[x]; ok {
						continue
					}
				} else {
					temp[x] = struct{}{}
					result = append(result, x)
				}

				for y = 0; y < set; y++ {

					if _, ok := temp[y]; ok {
						continue
					}

					if x == y {
						continue
					}

					if levels[x] > levels[y] {
						differ[y] = levels[x] - levels[y]
					} else {
						differ[y] = levels[y] - levels[x]
					}

				}

				keys := make([]uint8, 0, len(differ))

				for key := range differ {
					keys = append(keys, key)
				}

				sort.SliceStable(keys, func(i, j int) bool {

					return differ[keys[i]] < differ[keys[j]]

				})

				var next int
				var k int
				var s uint8

				for s = 0; s < set; s++ {

					for k = 0; k < len(keys); k++ {

						if k+1 < len(keys) {
							next = k + 1
						} else {
							break
						}

						if differ[keys[k]] == differ[keys[next]] {

							if keys[k] > keys[next] {

								keys[k], keys[next] = keys[next], keys[k]

								continue
							}
						}
					}
				}

				for k, val := range keys {

					if k == 0 {

						if x > 0 {
							temp[x] = struct{}{}
							result = append(result, x)
						}
						temp[val] = struct{}{}
						result = append(result, val)

					}
				}
			}

			result = append(result, 101)

		}

		if i%2 != 0 { //odd
			fmt.Fscan(reader, &set)
		}

		i++
	}

	var one uint8

	n := 0
	for _, v := range result {

		if v == 101 {
			one = 0
			n = 0
			fmt.Fprintf(output, " \n")
		} else {

			if n%2 != 0 {
				fmt.Fprintln(output, fmt.Sprintf("%v %v", one+1, v+1))
			}
			one = v
			n++
		}
	}

}
