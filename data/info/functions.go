package info

import (
	"context"
	"myproj/common"
	"time"
	"myproj/data/coal"
	"myproj/errors"
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

func AddMiner(m coal.MinerInterface) (int, error, *coal.MinerInfo) {
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
		
		return newMiner.Id, nil, &newMiner
	} else {
		return 0, errors.ErrMinerNotEnoughMoney, nil;
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