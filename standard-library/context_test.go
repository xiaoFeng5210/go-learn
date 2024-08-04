package standardLibrary

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background() // 起点
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()
	end := make(chan struct{}, 1)
	ch := timeoutCtx.Done()

	go func() {
		Sleep2s()
		end <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("end")
	}

}

func Sleep2s() {
	time.Sleep(2 * time.Second)
}
