package equipment

import (
	"encoding/json"
	"myproj/errors"
	"myproj/data/info"
	"net/http"
	"sync"
)

type Equipments struct {
	items map[string]*Equipment
	mtx   sync.RWMutex
}

func NewEquipments() *Equipments {
	e := &Equipments{
		items: make(map[string]*Equipment),
	}
	
	e.items["pickaxes"] = NewEquipment(3000)
	e.items["ventilation"] = NewEquipment(15000)
	e.items["trolleys"] = NewEquipment(50000)
	
	return e
}

func Сompletion() bool {
	equipments.mtx.RLock();
	defer equipments.mtx.RUnlock();

	for _, v := range equipments.items {
		if !v.IsPurchased() {
			return false;
		}
	}

	return true;
}

func GetPurchasedInfo(w http.ResponseWriter, r *http.Request) error {
	equipments.mtx.RLock()
	defer equipments.mtx.RUnlock()

	if err := json.NewEncoder(w).Encode(map[string]bool{
		"Pickaxe":     equipments.items["pickaxes"].IsPurchased(),
		"Ventilation": equipments.items["ventilation"].IsPurchased(),
		"Trolleys":    equipments.items["trolleys"].IsPurchased(),
	}); err != nil {
		return err
	}
	return nil
}

func BuyEquipment(name string) error {
	equipments.mtx.Lock();
	defer equipments.mtx.Unlock();

	eq, exists := equipments.items[name];
	if !exists {
		return errors.ErrEquipmentNotFound;
	}

	if eq.IsPurchased() {
		return errors.ErrEquipmentAlreadyPurchased;
	}

    currentWallet := info.MyCompany.GetWallet()
    if currentWallet < equipments.items[name].cost {
        return errors.ErrEquipmentNotEnoughMoney
    }

	info.MyCompany.SetWallet(-equipments.items[name].cost);

	eq.Complete();
	return nil;
}