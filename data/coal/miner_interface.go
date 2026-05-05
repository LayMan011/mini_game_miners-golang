package coal

import "context"

type MinerInterface interface {
	Run(ctx context.Context) <-chan int
	Info() MinerInfo
	GetClass() string
}