package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	degree    int
	min       int
	remainder = make([]int, 0)
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	CompetitiveResults(reader)
}

func CompetitiveResults(reader *bufio.Reader) {

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var sets int
	fmt.Fscan(reader, &sets)

	for i := 0; i < sets; i++ {

		var count int
		fmt.Fscan(reader, &count)

		set := make([]int, count, count)
		result := make([]int, count, count)

		for n := 0; n < count; n++ {
			fmt.Fscan(reader, &set[n])
		}

		for n := 1; n <= count; n++ {

			if n == 1 {
				degree = 1
			}

			if len(remainder) > 0 {
				min = remainder[0]
				for _, element := range remainder {

					if element < min {
						min = element
					}
				}

			} else {
				min = set[0]
				for k, element := range set {

					if result[k] != 0 {
						continue
					}

					if element < min {
						min = element
					}
				}
			}

			k := 0
			temp := make([]int, 0, count)

			for j := 0; j < count; j++ {

				remainder = make([]int, 0, count)

				for index := 0; index < count; index++ {

					if result[index] != 0 {
						continue
					}

					if set[index] == min || set[index] == min+1 {
						result[index] = degree
						temp = append(temp, set[index])
						k++
					} else {
						remainder = append(remainder, set[index])
					}
				}

				//Redefine min as max from temp.
				if len(temp) > 0 {

					min = temp[0]
					for _, element := range temp {
						if element > min {
							min = element
						}
					}

				}
			}

			degree = degree + k

		}

		for k, r := range result {

			fmt.Fprint(writer, r)

			if k == count-1 {
				fmt.Fprintln(writer, "")
			} else {
				fmt.Fprint(writer, " ")
			}
		}

	}

}
