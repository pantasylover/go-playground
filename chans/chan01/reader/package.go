package reader

import "fmt"

func StartReader(myChan chan string) {
	for {
		s := <-myChan

		fmt.Printf("[READER] Found string: %s\n", s)
	}
}
