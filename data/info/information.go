package info

import (
	"myproj/data/coal"
	"sync"
	"time"
)

type Company struct {
	MinersNow map[string](map[int]coal.MinerInfo)
	minersAll map[string](map[int]coal.MinerInfo)
	wallet int
	timeStart time.Time

	mtx sync.RWMutex
}

func NewCompany() *Company {
	return &Company{
		MinersNow: make(map[string]map[int]coal.MinerInfo),
		minersAll: make(map[string]map[int]coal.MinerInfo),
		wallet: 0,
		timeStart: time.Now(),
	}
}

func (c *Company) GetWallet() int {
    c.mtx.RLock()
    defer c.mtx.RUnlock()
    return c.wallet
}

func (c *Company) SetWallet(number int) {
	c.mtx.RLock()
    defer c.mtx.RUnlock()
    c.wallet += number;
}

func (c *Company) GetMinersNow() map[string](map[int]coal.MinerInfo) {
    c.mtx.RLock()
    defer c.mtx.RUnlock()

	result := make(map[string]map[int]coal.MinerInfo)
    for class, miners := range c.MinersNow {
        result[class] = make(map[int]coal.MinerInfo)
        for id, miner := range miners {
            result[class][id] = miner
        }
    }

    return result;
}

func (c *Company) GetMinersAll() map[string](map[int]coal.MinerInfo) {
    c.mtx.RLock()
    defer c.mtx.RUnlock()

    result := make(map[string]map[int]coal.MinerInfo)
    for class, miners := range c.minersAll {
        result[class] = make(map[int]coal.MinerInfo)
        for id, miner := range miners {
            result[class][id] = miner
        }
    }

    return result;
}

func (c *Company) GetTimeStart() time.Time {
    c.mtx.RLock()
    defer c.mtx.RUnlock()
    return c.timeStart;
}

func (c *Company) GetFullTime() time.Duration {
    c.mtx.RLock()
    defer c.mtx.RUnlock()
    return time.Since(c.timeStart);
}

func (c *Company) GetLenMinersNow() int {
    c.mtx.RLock()
    defer c.mtx.RUnlock()
    
    total := 0
    for _, miners := range c.MinersNow {
        total += len(miners)
    }
    return total
}

func (c *Company) GetLenMinersAll() int {
    c.mtx.RLock()
    defer c.mtx.RUnlock()
    
    total := 0
    for _, miners := range c.minersAll {
        total += len(miners)
    }
    return total
}