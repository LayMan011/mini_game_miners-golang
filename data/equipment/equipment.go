package equipment

import (
	"errors"
	"sync"
)

type Equipment struct {
	cost  int
	isBuy bool
	mtx   sync.RWMutex
}

var (
	classPickaxes = "pickaxes"
	classVentilation = "ventilation"
	classTrolleys    = "trolleys"
)

func NewEquipment(cost int) *Equipment {
	return &Equipment{
		cost: cost,
		isBuy: false,
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

	e.isBuy = true;
}

func (e *Equipment) IsPurchased() bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.isBuy;
}

var pickaxes = NewEquipment(3000);
var ventilation = NewEquipment(15000);
var trolleys = NewEquipment(50000);