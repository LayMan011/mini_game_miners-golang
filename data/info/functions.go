package info

import (
	"context"
	"myproj/common"
	"time"
)

func ReadFromCh(ch <-chan int) {
	for coal := range ch {
		MyCompany.mtx.Lock()
		MyCompany.wallet += coal
		MyCompany.mtx.Unlock()
	}
}

func PassiveIncome(ch chan<- int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case <-time.After(1 * time.Second):
			ch <- 1
		}
	}
}

func StartPassiveIncome() <-chan int {
	passiveCh := make(chan int)
	go PassiveIncome(passiveCh, common.BackCtx)
	return passiveCh
}