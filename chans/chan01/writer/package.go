package writer

import (
	"fmt"
	"time"
)

func StartWriter(myChan chan string) {
	for {
		str := fmt.Sprintf("[WRITER] Message at %d\n", time.Now().Unix())
		myChan <- str

		time.Sleep(time.Second * 1)
	}
}
