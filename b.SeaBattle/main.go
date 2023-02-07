package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	SatisfyCondition(reader)
}

func SatisfyCondition(reader *bufio.Reader) {

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var sets int
	fmt.Fscan(reader, &sets)

	for i := 0; i < sets; i++ {

		set := make([]int, 10)

		for n := 0; n < 10; n++ {
			fmt.Fscan(reader, &set[n])
		}

		type rules struct {
			one   int
			two   int
			three int
			four  int
		}

		var r = rules{}

		for _, v := range set {

			switch v {
			case 1:
				r.one++
			case 2:
				r.two++
			case 3:
				r.three++
			case 4:
				r.four++
			}
		}

		if r.one == 4 && r.two == 3 && r.three == 2 && r.four == 1 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}

	}
}
