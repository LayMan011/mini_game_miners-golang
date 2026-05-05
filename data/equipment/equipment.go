package equipment

import (
	"errors"
	"sync"
)

type Equipment struct {
	Cost  int
	IsBuy bool
	mtx   sync.RWMutex
}

func NewEquipment(cost int) *Equipment {
	return &Equipment{
		Cost: cost,
		IsBuy: false,
	}
}

func NewEquipmentsType(equipment string) (*Equipment, error) {
	switch equipment {
	case classPickaxes:
		return pickaxes, nil
	case classVentilation:
		return ventilation, nil
	case classTrolleys:
		return trolleys, nil
	default:
		return &Equipment{}, errors.New("there is no such equipment");
	}
}

func (e *Equipment) Complete() {
	e.mtx.Lock();
	defer e.mtx.Unlock();

	e.IsBuy = true;
}

func (e *Equipment) IsPurchased() bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.IsBuy;
}