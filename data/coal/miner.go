package coal

import (
	"context"
	"sync"
	"time"
)

var mtx sync.RWMutex

type Miner struct {
	class         string
	cost 		  int
	energy        int
	oneExtraction int
	breakMiner    int
	progress 	  int
}

func NewMiner(class string, cost int, energy int, one_extraction int, mybreak int, progress int) *Miner {
	return &Miner{
		class: class,
		cost: cost,
		energy: energy,
		oneExtraction: one_extraction,
		breakMiner: mybreak,
		progress: progress,
	};
}

func (m *Miner) Run(ctx context.Context) <-chan int  {
	transferPoint := make(chan int);

	go func() {
		defer close(transferPoint);

		for {
			select {
			case <-ctx.Done():
				return;
			default:
				time.Sleep(time.Duration(m.breakMiner) * time.Second);

                select {
                case <-ctx.Done():
                    return
                case transferPoint <- m.oneExtraction:
                }
		
				m.energy--;
				
				if m.progress != 0 {
					m.oneExtraction += m.progress;
				}
		
				if m.energy == 0 {
					return;
				}
			}
		}
	}()

	return transferPoint;
}

func (m *Miner) Info() MinerInfo {
	mtx.Lock();
	defer mtx.Unlock();

	var id int
	switch(m.class) {
	case classLittleMiner:
		idLittleMiner++;
		id = idLittleMiner;
	case classNormalMiner:
		idNormalMiner++;
		id = idNormalMiner
	case classBigMiner:
		idBigMiner++;
		id = idBigMiner
	}

	return MinerInfo{
		Id: id,
		Cost: m.cost,
		Class: m.class,
		Energy_remained: &m.energy,
	}
}

func (m *Miner) GetClass() string {
	mtx.RLock();
	defer mtx.RUnlock();
	return m.class;
}