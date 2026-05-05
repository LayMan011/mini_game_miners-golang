package info

import (
	"myproj/data/coal"
	"myproj/errors"
	"sync"
	"time"
)

var MyCompany = NewCompany();

type Company struct {
	MinersNow map[string](map[int]coal.MinerInfo)
	minersAll map[string](map[int]coal.MinerInfo)
	wallet int
	timeStart time.Time

	mtx sync.RWMutex
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

func NewCompany() *Company {
	return &Company{
		MinersNow: make(map[string]map[int]coal.MinerInfo),
		minersAll: make(map[string]map[int]coal.MinerInfo),
		wallet: 0,
		timeStart: time.Now(),
	}
}

func AddMiner(m coal.MinerInterface) (int, error) {
	MyCompany.mtx.Lock();
	defer MyCompany.mtx.Unlock();

	newMiner := m.Info();
	if MyCompany.wallet >= newMiner.GetCost() {
		MyCompany.wallet -= newMiner.GetCost();

		if MyCompany.minersAll[m.GetClass()] == nil {
			MyCompany.minersAll[m.GetClass()] = make(map[int]coal.MinerInfo);
		}

		if MyCompany.MinersNow[m.GetClass()] == nil {
			MyCompany.MinersNow[m.GetClass()] = make(map[int]coal.MinerInfo);
		}
		
		MyCompany.minersAll[m.GetClass()][newMiner.GetId()] = newMiner;
		MyCompany.MinersNow[m.GetClass()][newMiner.GetId()] = newMiner;
		
		return newMiner.Id, nil
	} else {
		return 0, errors.ErrMinerNotEnoughMoney;
	}
}

func DeleteMiner(class string, id int) error {
	if _, ok := MyCompany.MinersNow[class]; !ok {
		return errors.ErrClassMinerNotFound
	}

	if _, ok := MyCompany.MinersNow[class][id]; !ok {
		return errors.ErrMinerNotFound
	}

	delete(MyCompany.MinersNow[class], id)

	return nil
}

func GetMinersNow() map[string](map[int]coal.MinerInfo) {
	return MyCompany.MinersNow;
}

func GetMinersNowClass(class string) (map[int]coal.MinerInfo, error) {
	miners, ok := MyCompany.MinersNow[class];
	if !ok {
		return make(map[int]coal.MinerInfo), errors.ErrClassMinerNotFound;
	}

	return miners, nil;
}