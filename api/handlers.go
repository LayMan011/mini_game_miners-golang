package api

import (
	"context"
	"encoding/json"
	"fmt"
	"myproj/data/coal"
	"myproj/common"
	"myproj/data/equipment"
	"myproj/errors"
	"myproj/data/info"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	company *info.Company
	server *http.Server
}

func NewHTTPHandlers(company *info.Company) *HTTPHandlers {
	return &HTTPHandlers{
		company: company,
	}
}

/*
pattern: /miners/info
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent payments Miners
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerMinersInfo(w http.ResponseWriter, r *http.Request) {
	WriteData(map[string]int{
		"Little miner":   5,
		"Normal miner": 50,
		"Big miner": 450,
	}, w);
}

/*
pattern: /miners/{class}
method: POST
info: pattern

succed:
 - status code: 201 Created
 - response body: JSON represented MinerInfo
failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerMinerAdd(w http.ResponseWriter, r *http.Request) {
	class := mux.Vars(r)["class"];

	miner, err := coal.NewMinersType(class);
	if err != nil {
		errors.HttpErrorBadRequest(w, err);
		return;
	}

	id, err, newMinerInfo := info.AddMiner(miner);
	if err != nil {
		errors.HttpErrorBadRequest(w, err);
		return;
	}

	ChX := miner.Run(common.BackCtx);

    go func() {
        for value := range ChX {
            info.MyCompany.SetWallet(value)
        }

        if err := info.DeleteMiner(class, id); err != nil {
            errors.HttpErrorBadRequest(w, err);
        }
    }()

	WriteData(newMinerInfo, w);	
}

/*
pattern: /miners?now=true
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent Miners Now
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerMinersNow(w http.ResponseWriter, r *http.Request) {
	miners := info.GetMinersNow();

	WriteData(miners, w);
}

/*
pattern: /miners/{class}
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent Miners All
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerMinersNowClass(w http.ResponseWriter, r *http.Request) {
	class := mux.Vars(r)["class"];

	miners, err := info.GetMinersNowClass(class);
	if err != nil {
		errors.HttpErrorBadRequest(w, err);
		return;
	}

	WriteData(miners, w);
}

/*
pattern: /equipments/info
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent payments Equipmetns
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerPriceEquipmetns(w http.ResponseWriter, r *http.Request) {
	WriteData(map[string]int{
		"Pickaxe":     3_000,
		"Ventilation": 15_000,
		"Trolleys":    50_000,
	}, w);
}

/*
pattern: /equipments/{equipment}
method: PATCH
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represented purchased equipment
failed:
 - status code: 400, 404, 409, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerBuyEquipments(w http.ResponseWriter, r *http.Request) {
	class := mux.Vars(r)["equipment"];

	NewEquipment, err := equipment.NewEquipmentsType(class);
	if err != nil {
		errors.HttpErrorBadRequest(w, err);
		return;
	}

	if NewEquipment.IsPurchased() {
		errors.HttpErrorBadRequest(w, fmt.Errorf("equipment already purchased"))
        return
	}

	if err := equipment.BuyEquipment(class); err != nil {
		errors.HttpErrorEquipmentBuy(w, err);
		return;
	}

	NewEquipment.IsBuy = true;

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(NewEquipment); err != nil {
		panic(err)
	}
}

/*
pattern: /equipments
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent payments Miners
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerCompleteEquipments(w http.ResponseWriter, r *http.Request) {
	equipment.GetPurchasedInfo(w, r);
}

/*
pattern: /info
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent payments Miners
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerInfoCompany(w http.ResponseWriter, r *http.Request) {
	WriteData(map[string]int{
		"Wallet":     info.MyCompany.GetWallet(),
		"LenMinersNow": info.MyCompany.GetLenMinersNow(),
		"LenMinersAll":    info.MyCompany.GetLenMinersAll(),
	}, w);
}

/*
pattern: /end
method: GET
info: -

succed:
 - status code: 200 Ok
 - response body: JSON represent payments Miners
failed:
 - status code: 400, 500, ... 
 - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerEndGame(w http.ResponseWriter, r *http.Request) {
	if equipment.Сompletion() {
		WriteData(map[string]any{
			"MinersNow": info.MyCompany.GetMinersNow(),
			"MinersAll": info.MyCompany.GetMinersAll(),
			"Wallet": info.MyCompany.GetWallet(),
			"TimeStart": info.MyCompany.GetTimeStart(),
			"FullTime": fmt.Sprintf("%.2f minutes", info.MyCompany.GetFullTime().Minutes()),
		}, w);

		common.BackCtxCancel();
		go func() {
			_ = h.server.Shutdown(context.Background())
		}()
	} else {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode("You haven't purchased all the equipment, so the game can't be completed!"); err != nil {
			panic(err);
		}
	}
}